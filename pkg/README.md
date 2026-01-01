# pkg/

Public library code that can be imported by external projects.

## Structure

- `k8s/` - Kubernetes client utilities
- `mps/` - MPS GPU sharing management
- `logger/` - Logging utilities

## Purpose

Contains reusable utilities and helpers that could be shared with other projects. These are infrastructure-level utilities, not business logic.

## MPS GPU Sharing

MPS (Multi-Process Service) enables efficient GPU sharing:
- 1 dedicated GPU = 10 MPS units
- Project-level MPS limits
- Validation against quotas
- Dynamic allocation tracking
