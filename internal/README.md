# internal/

Private application code that cannot be imported by external projects.

## Structure

- `api/` - HTTP API layer (routes, handlers, middleware)
- `domain/` - Business entities and domain logic
- `scheduler/` - Job scheduling implementation
- `priority/` - Resource priority and preemption logic
- `repository/` - Data access layer
- `service/` - Business logic layer

## Purpose

Contains all core business logic organized by feature and layer. This code is private to the platform-go project.
