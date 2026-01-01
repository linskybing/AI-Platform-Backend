# internal/api/

HTTP API layer for the platform.

## Structure

- `routes/` - Route definitions and grouping
- `handlers/` - HTTP request handlers
- `middleware/` - Authentication, logging, error handling

## Guidelines

- Handlers should be < 100 lines
- Use dependency injection for services
- Follow RESTful conventions
- All routes versioned under `/api/v1/`
- Include Swagger annotations for API documentation
