# Learning Center Management System

A Go-based backend system for managing learning centers, courses, students, instructors, and related operations.

## Project Structure

```
.
├── cmd/                  # Application entry points
│   └── api/              # Main API server
├── internal/             # Private application code
│   ├── auth/             # Authentication and authorization
│   ├── handlers/         # HTTP request handlers
│   ├── middleware/       # HTTP middleware components
│   ├── models/           # Data models and business objects
│   ├── repository/       # Data access layer
│   └── services/         # Business logic layer
├── pkg/                  # Public libraries that can be used by external applications
│   ├── config/           # Configuration management
│   ├── database/         # Database utilities
│   ├── logger/           # Logging utilities
│   └── validator/        # Validation utilities
├── scripts/              # Scripts for development, CI/CD, etc.
│   └── migrations/       # Database migration scripts
└── docs/                 # Documentation
```

## Layer Architecture

This project follows a layered architecture:

1. **Handlers Layer** - Handles HTTP requests and responses
2. **Service Layer** - Contains business logic
3. **Repository Layer** - Manages data access and persistence
4. **Model Layer** - Defines data structures

## Getting Started

[Instructions for setting up and running the project will be added here]

## API Documentation

[API documentation will be added here]