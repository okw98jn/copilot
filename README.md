# Copilot Project

This repository contains a Docker-based development environment for a web application with Go backend API and React frontend, using PostgreSQL as the database.

## Docker Environment Setup

### Prerequisites

- Docker and Docker Compose installed on your system
- Git (to clone the repository)

### Services

- **PostgreSQL**: Database server
- **Go API**: Backend API service (placeholder)
- **React**: Frontend service (placeholder)

### Setup Instructions

1. Clone the repository:

```bash
git clone https://github.com/okw98jn/copilot.git
cd copilot
```

2. Start the Docker environment:

```bash
docker compose up -d
```

3. Access the services:

- PostgreSQL: localhost:5432
- Go API: http://localhost:8080
- React Frontend: http://localhost:3000

### Container Management

- To stop all containers:

```bash
docker compose stop
```

- To stop and remove all containers:

```bash
docker compose down
```

- To stop and remove all containers including volumes (will delete database data):

```bash
docker compose down -v
```

- To view container logs:

```bash
docker compose logs -f [service_name]
```

### Database Connection

- Host: localhost
- Port: 5432
- Username: postgres
- Password: postgres
- Database: copilot_db

## Development

- Backend code is in the `backend/` directory
- Frontend code is in the `frontend/` directory
- Database initialization scripts can be placed in `db/init/`
