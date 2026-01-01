# Platform-Go Project Structure Guide

## Current Architecture Overview

Platform-Go is organized with **hexagonal/layered architecture** with clear separation between:
- **Entry Points** (`cmd/`)
- **Domain Logic** (`internal/domain/`)
- **Application Services** (`internal/application/`)
- **Infrastructure** (`internal/repository/`, `infra/`)
- **Utilities & Libraries** (`pkg/`)

## Directory Organization

### `cmd/` - Application Entry Points
Two independent executable binaries:

```
cmd/
 api/ # HTTP REST API server
 main.go # Entry point (port 8080)
 scheduler/ # Background job scheduler
 main.go # Entry point (standalone process)
```

**Key Points:**
- API and Scheduler are completely decoupled
- Each has its own `main.go` and can be built/deployed independently
- No cross-imports between them

### `internal/` - Core Application Logic

#### `domain/` - Entity & Business Models
Contains domain entities and their transfer objects (DTOs):

```
internal/domain/
 audit/ # Audit trail entities
 configfile/ # Config file entities
 course/ # Course entities
 form/ # Form submission entities
 gpu/ # GPU resource entities
 group/ # User group entities
 job/ # Job entities
 project/ # Project entities
 resource/ # K8s resource entities
 user/ # User entities
 view/ # Query result view models (read-only DTOs)
```

Each domain package typically contains:
- `model.go` - Core entity definitions (for database)
- `dto.go` - Input/Output transfer objects

**The `view/` Package:**
- Contains read-only query results
- Used for API responses that combine multiple entities
- Example: `UserGroupView` has Username, GroupName in addition to UID, GID

#### `application/` - Business Logic Services
High-level business operations and orchestration:

```
internal/application/
 audit_service.go
 audit_service_test.go
 configfile_service.go
 configfile_service_test.go
 group_service.go
 group_service_test.go
 job/ # Job-related business logic
 project_service.go
 project_service_test.go
 resource_service.go
 resource_service_test.go
 scheduler/ # Scheduler service
 user_service.go
 user_service_test.go
 usergroup_service.go
 usergroup_service_test.go
```

**Purpose:**
- Implements use cases (e.g., CreateProject, DeleteConfigFile)
- Coordinates between repositories and domain logic
- Handles transactions and validation

#### `api/` - HTTP Interface Layer
REST API request handlers:

```
internal/api/
 handlers/ # HTTP request handlers
 middleware/ # HTTP middleware (CORS, logging, auth)
 routes/ # Route definitions
```

#### `repository/` - Data Access Layer
Abstractions for database operations:

```
internal/repository/
 mock/ # Mock implementations for testing
 audit.go
 configfile.go
 group.go
 job.go
 project.go
 resource.go
 user.go
 usergroup.go
```

**Pattern:**
- Each domain entity has a corresponding repository interface
- `DBXxxRepo` implements the actual database access
- `MockXxxRepo` for unit testing

#### `scheduler/` - Job Scheduling Infrastructure
Background job processing system:

```
internal/scheduler/
 executor/ # Job executor registry and implementations
 gpu/ # GPU-specific scheduling logic
 mpi/ # MPI job scheduling
 mps/ # Multi-Process Service handling
 queue/ # Job queue implementation
```

#### `priority/` - Resource Priority Management
Priority-based resource allocation:

```
internal/priority/
 monitor/ # Resource monitoring (CPU, memory, GPU)
 preemptor/ # Job preemption logic
```

#### `config/` - Configuration Management
```
internal/config/
 db/ # Database configuration
```

### `pkg/` - Public Utilities & Libraries
Reusable packages that could be used by external projects:

```
pkg/
 k8s/ # Kubernetes client utilities
 logger/ # Logging utilities
 mps/ # Multi-Process Service utilities
 response/ # HTTP response formatters
 storage/ # Storage (MinIO) utilities
 types/ # Common type definitions
 utils/ # General utility functions (YAML, JSON parsing)
```

### `infra/` - Infrastructure Configuration
Deployment and infrastructure setup:

```
infra/
 db/ # Database initialization
 postgres/
 minio/ # MinIO object storage config
 values.yaml # Helm values
```

### `k8s/` - Kubernetes Manifests
Declarative infrastructure-as-code:

```
k8s/
 ca.yaml # Certificate Authority
 go-api.yaml # API Deployment & Service
 go-scheduler.yaml # Scheduler Deployment (NEW)
 postgres.yaml # PostgreSQL StatefulSet
 secret.yaml # K8s Secrets
 storage.yaml # Storage classes & PVCs
```

### `scripts/` - Helper Scripts
Build and deployment automation:

```
scripts/
 build_images.sh # Build Docker images
 clean_dev.sh # Clean development environment
 create_gpu_pod.py # GPU pod creation utility
 dev.sh # Development setup
 fix.sh # Code fixing utilities
 genmock.sh # Generate mock files
 setup_fakek8s.sh # Local K8s setup
```

## Recommended Organization Improvements

### 1. Create Documentation
**Done:** Created this file at `docs/PROJECT_STRUCTURE.md`

### 2. Separate K8s Deployments
**Done:** Created `k8s/go-scheduler.yaml` with:
- Independent Scheduler Deployment
- ServiceAccount with proper RBAC
- Resource requests/limits
- Health checks (liveness/readiness probes)

### 3. Improve Internal Structure (Optional but Recommended)

Currently `internal/` has both `application/` and `scheduler/`:
- `internal/application/scheduler/` - Business service
- `internal/scheduler/` - Infrastructure

**Suggestion:** Could add `internal/service/` layer if business logic grows:

```
internal/service/ # Domain-specific business logic
 audit/
 configfile/
 group/
 project/
 resource/
```

Then `internal/application/` becomes:
```
internal/application/ # High-level orchestration
 use_cases/
 job/ # Job-specific orchestration
```

### 4. Enhance Deployment Strategy

**Current:** Single `deploy/` directory (if exists)

**Recommended:**
```
deploy/
 docker/
 api.Dockerfile
 scheduler.Dockerfile
 kubernetes/
 base/
 api/
 scheduler/
 overlays/
 dev/
 staging/
 production/
 helm/
 platform-go/
 Chart.yaml
 values.yaml
 templates/
 README.md
```

## Testing Strategy

### Unit Tests
```bash
go test ./... -v -short
```
- Tests in `*_test.go` files alongside source code
- Uses mock repositories via `internal/repository/mock/`
- Fast execution (< 1 second)

### Integration Tests
```bash
go test ./... -v -tags=integration
```
- Requires PostgreSQL running
- Tests full stack with real database
- Defined in `.github/workflows/integration-test.yml`

### Build Verification
```bash
make build # Build both API & Scheduler
make build-api # Build API only
make build-scheduler # Build Scheduler only
```

## CI/CD Pipeline

### GitHub Actions Workflow: `integration-test.yml`

**Triggers:**
- Push to `main` branch
- Pull requests
- Manual `workflow_dispatch`
- Changes to Go files, `go.mod`, `go.sum`, `deploy/**`, `k8s/**`

**Pipeline Steps:**
1. Checkout code
2. Setup Go 1.22
3. Download dependencies
4. Install kubectl
5. Setup Kind (local K8s cluster)
6. Create test cluster
7. Wait for PostgreSQL
8. Run `go test ./... -v -tags=integration`
9. Collect results

## Best Practices

### Layer Responsibilities

| Layer | Responsibility | Example |
|-------|---|---|
| **API** | HTTP protocol handling, request/response formatting | Parse JSON, validate headers |
| **Application** | Business logic orchestration | DeleteUserGroup validates permissions |
| **Domain** | Core entities and value objects | UserGroup struct with validation |
| **Repository** | Data persistence abstraction | Query database, return entities |
| **Infrastructure** | External services, K8s operations | Create namespace, apply RBAC |

### Import Flow (Allowed)
```
cmd → internal/application → internal/domain
 ↓
internal/repository → internal/domain
↓
internal/api → internal/application
```

### Import Flow (NOT Allowed)
```
internal/domain → internal/application 
internal/api → internal/domain 
pkg → internal 
```

## Code Organization Checklist

- [x] Entry points separated (`cmd/api`, `cmd/scheduler`)
- [x] Clear domain models in `internal/domain/`
- [x] Business logic in `internal/application/`
- [x] Data access abstracted in `internal/repository/`
- [x] Infrastructure code in `internal/scheduler/`, `infra/`
- [x] Public utilities in `pkg/`
- [x] K8s manifests under `k8s/`
- [x] Tests alongside source code (`*_test.go`)
- [x] Mock implementations in `internal/repository/mock/`
- [x] CI/CD pipeline configured in `.github/workflows/`
- [x] Separate K8s deployments for API and Scheduler

## Quick Reference

**Project Root:**
```
platform-go/
 cmd/ ← Application entry points (API & Scheduler)
 internal/ ← Core business logic (NOT exported)
 pkg/ ← Shared libraries (exportable)
 k8s/ ← Kubernetes manifests
 infra/ ← Infrastructure setup
 scripts/ ← Helper scripts
 deploy/ ← Docker configs
 docs/ ← Documentation
 Makefile ← Build targets
 go.mod ← Dependencies
 README.md ← Project overview
```

## Next Steps

1. Run full test suite: `go test ./...`
2. Build images: `make build`
3. Deploy to K8s:
 ```bash
 kubectl apply -f k8s/postgres.yaml
 kubectl apply -f k8s/go-api.yaml
 kubectl apply -f k8s/go-scheduler.yaml
 ```
4. Verify both services are running:
 ```bash
 kubectl get deployments
 kubectl logs -f deployment/go-api
 kubectl logs -f deployment/go-scheduler
 ```
