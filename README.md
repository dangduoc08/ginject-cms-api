# ginject-cms-api

Golang project with a Makefile supporting build, run, test, lint, and Docker/Docker Compose.

## Project Structure

- `cmd/app/` - Directory containing `main.go`
- `internal/` - Internal packages and business logic (not exposed outside the project)
- `bin/` - Directory for the built binary
- `deploy/docker/` - Dockerfile and docker-compose files
- `.env.example` - Example environment file

---

## Prerequisites

- Go >= 1.25
- Docker & Docker Compose
- Make
- golangci-lint (for lint command)

---

## Makefile Commands

You can list all Makefile commands:

```bash
make help
