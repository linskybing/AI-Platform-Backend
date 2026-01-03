# Kubernetes Client Initialization Fix

## Issue
Runtime panic in WebSocket storage status handler when attempting to list Kubernetes resources.

```
panic: runtime error: invalid memory address or nil pointer dereference
k8s.io/client-go/dynamic.(*dynamicResourceClient).list(0xc0002b8550, ...)
  /root/go/pkg/mod/k8s.io/client-go@v0.33.3/dynamic/simple.go:276 +0xd8
```

## Root Cause
The Kubernetes client package exposes a global `DynamicClient` variable that must be initialized via `k8s.Init()` before use. However:

- **cmd/api/main.go** - Never called `k8s.Init()`
- **cmd/scheduler/main.go** - Never called `k8s.Init()`

When the WebSocket handler attempted to call `k8s.WatchNamespaceResources()`, which uses `DynamicClient`, it was nil because the initialization function was never invoked.

## Solution
Added `k8s.Init()` call to both entry points, after database initialization:

### cmd/api/main.go
```go
func main() {
    config.LoadConfig()   // Load configuration
    db.Init()             // Initialize database
    k8s.Init()            // Initialize Kubernetes client ← NEW
    
    // ... rest of initialization
}
```

### cmd/scheduler/main.go
```go
func main() {
    config.LoadConfig()   // Load configuration
    db.Init()             // Initialize database
    k8s.Init()            // Initialize Kubernetes client ← NEW
    
    // ... rest of initialization
}
```

## What k8s.Init() Does
The initialization function in `pkg/k8s/client.go` performs:

1. **Load Kubeconfig** - Reads from KUBECONFIG env var, in-cluster config, or ~/.kube/config
2. **Create REST Config** - Establishes Kubernetes API configuration
3. **Create Clientset** - Creates standard Kubernetes client interface
4. **Create Discovery Client** - Retrieves available API resources
5. **Create REST Mapper** - Maps API groups/versions to REST resources
6. **Create Dynamic Client** - Initializes the dynamic resource client used for watching resources
7. **Configure Rate Limiting** - Sets QPS=50, Burst=100 to prevent client rate limiting

## Affected Functionality
- ✅ WebSocket storage status monitoring (`/k8s/users/:username/storage/status`)
- ✅ Kubernetes resource watching (pods, services, deployments)
- ✅ Pod terminal execution via WebSocket

## Testing
All initialization and integration verified:
- ✅ Code compiles without errors
- ✅ All unit tests pass (100+ tests)
- ✅ Code formatting passes (`make fmt-check`)
- ✅ Static analysis passes (`make vet`)
- ✅ Both binaries build successfully

## Initialization Order
Critical initialization sequence in both main.go files:
1. `config.LoadConfig()` - Must be first (loads all configuration)
2. `db.Init()` - Must be before database usage
3. `k8s.Init()` - Must be before Kubernetes operations
4. Database AutoMigrate - Creates tables
5. Application startup (Router or Scheduler)
