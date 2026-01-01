# Platform-Go Quick Reference Card

## 

### 
```bash
make test # 
make test-unit # 
make test-verbose # 
make test-coverage # 
make test-race # 
make coverage-html # HTML 
```

### 
```bash
make fmt # 
make fmt-check # 
make vet # 
```

### 
```bash
make build # API Scheduler
make build-api # API
make build-scheduler # Scheduler
make clean # 
```

### Kubernetes
```bash
make k8s-deploy # K8s
make k8s-delete # K8s 
make k8s-status # 
make k8s-logs-api # API 
make k8s-logs-scheduler # Scheduler 
```

### CI/CD
```bash
make ci # CI 
make local-test # + 
make all # + 
```

---

## 

| | |
|---|---|
| `cmd/api/` | API |
| `cmd/scheduler/` | |
| `internal/api/` | HTTP handlers, routes, middleware |
| `internal/application/` | |
| `internal/domain/` | |
| `internal/repository/` | |
| `internal/scheduler/` | |
| `internal/priority/` | |
| `pkg/` | |
| `k8s/` | Kubernetes |
| `infra/` | |
| `docs/` | |

---

## 

| | |
|---|---|
| `go.mod` / `go.sum` | Go |
| `Makefile` | |
| `README.md` | |
| `docs/PROJECT_STRUCTURE.md` | |
| `docs/TESTING_REPORT.md` | |
| `.github/workflows/integration-test.yml` | CI |

---

## 

 
- `internal/application` (50+ )
- `internal/application/scheduler` (10+ )
- `internal/priority` (3 )
- `internal/priority/monitor` (4 )
- `internal/scheduler/executor` (8 )
- `internal/scheduler/queue` (2 )
- `pkg/mps` (4 )
- `pkg/utils` (8 )

 :
- `cmd/api/`, `cmd/scheduler/` - 
- `internal/api/handlers/`, `internal/api/middleware/`, `internal/api/routes/`
- `internal/domain/*` - 
- `internal/repository/` - 
- 

---

## 

### 
```bash
go test ./internal/application -v
go test ./internal/scheduler/executor -v
```

### 
```bash
go test ./internal/application -run TestCreateUserGroup -v
```

### 
```bash
go test ./... -coverprofile=coverage.out
go tool cover -func=coverage.out
go tool cover -html=coverage.out # 
```

### 
```bash
go test ./... -race
```

---

## 

### Q: Kubernetes
```bash
make k8s-deploy
```

### Q: Scheduler
```bash
make build-scheduler
```

### Q: 
```bash
make k8s-logs-api
make k8s-logs-scheduler
```

### Q: 
```bash
make fmt # 
make fmt-check # 
```

### Q: CI 
```bash
make fmt-check && make vet && make test && make build
```

---

## 

```bash
./platform-api # REST API (~72MB)
./platform-scheduler # (~2.4MB)
```

---

## 

Kubernetes 
- `k8s/secret.yaml` - 

1. PostgreSQL 
2. 
3. 

---

## 

```bash
DB_HOST=postgres
DB_PORT=5432
DB_USER=<from-secret>
DB_PASSWORD=<from-secret>
LOG_LEVEL=info
```

---

## K8s

- **API**: 
 - : CPU 100m, 128Mi
 - : CPU 500m, 512Mi
 
- **Scheduler**:
 - : CPU 100m, 128Mi
 - : CPU 500m, 512Mi

---

## 

- API : `8080` → K8s NodePort `30080`
- : `5432` ()
- MinIO: `9000` ()

---

## 

| | | |
|---|---|---|
| 2026-01-01 | K8s Scheduler | `k8s/go-scheduler.yaml` |
| 2026-01-01 | | `docs/PROJECT_STRUCTURE.md` |
| 2026-01-01 | | `docs/TESTING_REPORT.md` |
| 2026-01-01 | Makefile | `Makefile` |
| 2026-01-01 | | `internal/application/*_test.go` |
| 2026-01-01 | | 12 |

---

*: 2026-01-01*
