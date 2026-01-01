# pkg/mps/

MPS (Multi-Process Service) GPU sharing management.

## Purpose

Provides utilities for managing GPU sharing using NVIDIA MPS. Enables efficient GPU resource allocation across multiple projects.

## MPS Unit Conversion

- 1 dedicated GPU = 10 MPS units
- Projects have configurable MPS limits
- Validation against project quotas
- Dynamic allocation tracking

## Components

- MPS limit validation
- Unit conversion utilities
- Project quota checking
- Resource accounting

## Example Usage

```go
// Validate MPS request
valid := mps.ValidateMPSRequest(projectID, requestedMPS)

// Convert dedicated GPUs to MPS units
mpsUnits := mps.ConvertToMPSUnits(dedicatedGPUs)

// Check project quota
available := mps.CheckAvailableMPS(projectID)
```
