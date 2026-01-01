# Platform API (Go)

RESTful API and Job Scheduler written in Go. Uses Gin (HTTP), PostgreSQL (database), and MinIO (storage). Includes full Kubernetes manifests and CI/CD pipeline.

**Status**: Production Ready

## Documentation

- **[Project Structure Guide](docs/PROJECT_STRUCTURE.md)** - Detailed architecture overview
- **[Testing Report](docs/TESTING_REPORT.md)** - Complete test results and analysis
- **[Quick Reference](docs/QUICK_REFERENCE.md)** - Common commands and tips

## Quick Start

### Prerequisites
- Go 1.21+
- PostgreSQL 12+
- (Optional) MinIO for object storage

### Local Development

```bash
# Download dependencies
go mod download

# Run all tests
go test ./...

# Build binaries
make build

# Build API only
make build-api

# Build Scheduler only
make build-scheduler
```

### Kubernetes Deployment

```bash
# Deploy all resources
make k8s-deploy

# View status
make k8s-status

# View logs
make k8s-logs-api
make k8s-logs-scheduler

# Cleanup
make k8s-delete
```

## Project Structure

```
platform-go/
 cmd/ # Application entry points
 api/ # REST API server
 scheduler/ # Background job scheduler
 internal/
 api/ # HTTP layer
 application/ # Business logic
 domain/ # Entity models
 repository/ # Data access
 scheduler/ # Scheduling infrastructure
 priority/ # Resource priority management
 pkg/ # Reusable packages
 k8s/ # Kubernetes manifests
 infra/ # Infrastructure setup
 docs/ # Documentation
```

## Common Commands

### Testing
```bash
make test # Run all tests
make test-coverage # Generate coverage report
make test-race # Check for race conditions
```

### Code Quality
```bash
make fmt # Format code
make fmt-check # Check format
make vet # Static analysis
```

### Build & Deploy
```bash
make build # Build API + Scheduler
make clean # Clean artifacts
make ci # Full CI pipeline
make all # Build + Deploy
```

See [Quick Reference](docs/QUICK_REFERENCE.md) for more commands.

## Test Results

- **Unit Tests**: 100+ tests, all passing
- **Code Format**: All files properly formatted
- **Static Analysis**: No issues found
- **Build**: Both API and Scheduler binaries successful

Test coverage:
- `internal/application` - 50+ tests 
- `internal/scheduler` - 10+ tests 
- `internal/priority` - 3 tests 
- `pkg/mps` - 4 tests 
- `pkg/utils` - 8 tests 

See [Testing Report](docs/TESTING_REPORT.md) for details.

## Docker & Kubernetes

### Build Docker Images
```bash
docker build -t platform-go-api:latest -f deploy/api.Dockerfile .
docker build -t platform-go-scheduler:latest -f deploy/scheduler.Dockerfile .
```

### Kubernetes Deployment
```bash
# Apply all manifests
kubectl apply -f k8s/

# Or apply individually
kubectl apply -f k8s/secret.yaml
kubectl apply -f k8s/postgres.yaml
kubectl apply -f k8s/go-api.yaml
kubectl apply -f k8s/go-scheduler.yaml
```

## Configuration

Environment variables (via K8s secrets):
```bash
DB_HOST=postgres
DB_PORT=5432
DB_USER=<from-secret>
DB_PASSWORD=<from-secret>
```

See `k8s/secret.yaml` for sensitive data management.

## Helper Scripts

- `scripts/dev.sh` - Local development setup
- `scripts/build_images.sh` - Docker build automation
- `scripts/setup_fakek8s.sh` - Local K8s cluster setup
- `scripts/create_gpu_pod.py` - GPU pod creation helper

## API Documentation

Swagger documentation available at `/swagger.html` when server is running.

See `internal/api/handlers/` for endpoint implementations.

## Development Tips

### Run specific test
```bash
go test ./internal/application -run TestCreateUserGroup -v
```

### Generate coverage HTML
```bash
make coverage-html # Opens in browser
```

### Check code format issues
```bash
make fmt-check
```

### View detailed test output
```bash
make test-verbose
```

## Architecture

**Two Independent Services:**
1. **API Server** (`cmd/api/`) - HTTP REST interface
 - Port 8080
 - Request/response handling
 - Business logic orchestration

2. **Scheduler** (`cmd/scheduler/`) - Background job processor
 - No HTTP interface
 - Queue-based job processing
 - Priority-based resource allocation

Both services share:
- Domain models (`internal/domain/`)
- Business logic (`internal/application/`)
- Data access layer (`internal/repository/`)

## Key Files

| File | Purpose |
|---|---|
| `Makefile` | Build and development tasks |
| `.github/workflows/integration-test.yml` | CI/CD pipeline |
| `go.mod` / `go.sum` | Go dependencies |
| `k8s/go-api.yaml` | API deployment config |
| `k8s/go-scheduler.yaml` | Scheduler deployment config |

## Features

- Modular architecture with clear separation of concerns
- Comprehensive unit tests (100+ tests)
- Kubernetes-ready with full YAML configs
- Independent API and Scheduler services
- Priority-based job scheduling
- GPU resource management
- Full audit logging
- Project and configuration file management

## Notes

- Read each helper script before running; some modify system state
- Ensure PostgreSQL is running before deployment
- MinIO setup is optional for object storage
- For detailed architecture, see [Project Structure Guide](docs/PROJECT_STRUCTURE.md)

## Contributing

1. Make changes to code
2. Run tests: `make test`
3. Format code: `make fmt`
4. Build: `make build`
5. Test in Kubernetes: `make k8s-deploy`

---

**Last Updated**: 2026-01-01 
**Status**: Production Ready

