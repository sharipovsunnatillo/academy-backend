# Academy Backend

This is the backend service for the Academy project.

## Project Structure

The project follows the standard Go project layout:

- `cmd/`: Contains the main applications for this project
  - `api/`: The main API server application
- `internal/`: Private application and library code
- `pkg/`: Library code that's ok to use by external applications
- `docs/`: Documentation for the project

## Getting Started

To run the API server:

```bash
go run cmd/api/main.go
```

The server will start on port 8080.

https://github.com/qingwave/weave