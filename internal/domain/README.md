# internal/domain/

Domain entities and business logic organized by aggregate.

## Structure

- `job/` - Job entities, types, and domain logic
- `course/` - Course/pod entities and management
- `resource/` - Resource entities and quotas
- `user/` - User entities and authentication
- `project/` - Project entities and membership

## Guidelines

- Each domain folder contains:
  - `model.go` - Entity definitions
  - `types.go` - Enums and type definitions
  - `repository.go` - Repository interface
  - `service.go` - Domain service interface
- Keep files < 100 lines
- English comments only
- Pure domain logic, no infrastructure dependencies
