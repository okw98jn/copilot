# Todo API with Clean Architecture

This is a Todo API implementation using clean architecture principles, with the following technologies:

- Gin Web Framework
- Bun ORM
- Wire for Dependency Injection
- golang-migrate for database migrations

## Architecture

The application follows Clean Architecture principles:

- **Domain Layer**: Contains the core business logic and entities
- **Use Case Layer**: Implements the application's use cases
- **Interface Layer**: Adapts the domain and use cases to external interfaces
- **Infrastructure Layer**: Contains external frameworks and tools

## Setup and Running

1. Setup the environment
```bash
make setup
```

2. Build and start the containers
```bash
make up
```

3. Access the API
- Todo List: GET http://localhost:8080/api/todos

## Development

To rebuild the containers:
```bash
make rebuild
```

To view logs:
```bash
make logs
```

To access the shell:
```bash
make shell-go
```
