# cmd/

Application entry points.

## Structure

- `api/` - HTTP API server entry point
- `scheduler/` - Job scheduler service entry point

## Purpose

Contains main applications for the platform. Each subdirectory has its own `main.go` that serves as an executable entry point.

- **api**: Starts the HTTP server with REST API endpoints
- **scheduler**: Runs the job scheduling service with priority-based resource allocation
