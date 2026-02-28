
# Events Booking APIs — Build Requirements

> **Quick Navigation:**
> - [Scope](#scope)
> - [Goals](#goals)
> - [Prerequisites](#prerequisites)
> - [Repository Layout](#repository-layout)
> - [Project Contract](#project-contract)
> - [API Summary](#api-summary)
> - [Event Model](#event-model)
> - [Environment Variables](#environment-variables)
> - [Local Build & Run](#local-build--run)
> - [Testing](#testing)
> - [Observability & Logging](#observability--logging)
> - [Authentication & Security](#authentication--security)
> - [Persistence & Data](#persistence--data)
> - [Dockerfile Example](#dockerfile-example)
> - [CI/CD Recommendations](#cicd-recommendations)
> - [Developer Guidance](#developer-guidance)
> - [Edge Cases & Considerations](#edge-cases--considerations)

---

## Scope

- **Backend REST API** for creating, updating, and booking events
- **Authentication** using JWT for protected endpoints
- **Minimal persistence** (file or lightweight DB) for demo; production should use a managed relational DB

---

## Goals

- Provide a reliable, well-documented developer workflow to build, run, and test the service locally
- Make the code easy to read and extend; enforce simple standards for configuration, logging, and error handling

---

## Prerequisites

- Go 1.20+ installed and available in PATH. Verify with:
    ```sh
    go version
    ```
- Git for source control
- (Optional but recommended) Docker 20+ for container builds
- A code editor (VS Code recommended) with Go extensions

---

## Repository Layout

- `cmd/` — application entrypoints (if present)
- `internal/` or `pkg/` — core application logic and reusable packages
- `api/` or `routes/` — HTTP route wiring and handlers
- `models/` — domain models (Event, User, Registration)
- `configs/` — configuration and environment handling
- `scripts/` — helper scripts for local development or migrations

---

## Project Contract

- **Input:** HTTP requests as defined in the API section below
- **Output:** JSON responses and HTTP status codes; structured logs for observability
- **Error modes:**
    - Validation errors (400)
    - Authentication/authorization (401/403)
    - Not-found (404)
    - Internal errors (500)

---

## API Summary

| Method | Endpoint                        | Description                                 | Auth Required |
|--------|----------------------------------|---------------------------------------------|--------------|
| GET    | /events                         | List all events                             | No           |
| GET    | /events/:id                     | Get event by id                             | No           |
| POST   | /events                         | Create a new event                          | Yes          |
| PUT    | /events/:id                     | Update an event (creator only)              | Yes          |
| DELETE | /events/:id                     | Delete an event (creator only)              | Yes          |
| POST   | /signup                         | Register a new user                         | No           |
| POST   | /login                          | Authenticate and receive a JWT              | No           |
| POST   | /events/:id/register            | Register the authenticated user for an event| Yes          |
| DELETE | /events/:id/register            | Remove the authenticated user's registration| Yes          |

---

## Event Model

| Field         | Type      | Description                        |
|---------------|-----------|------------------------------------|
| id            | int/UUID  | Unique event identifier            |
| name          | string    | Event name                         |
| description   | string    | Event description                  |
| location      | string    | Event location                     |
| starts_at     | RFC3339   | Start time (UTC)                   |
| ends_at       | RFC3339   | End time (UTC, optional)           |
| host_user_id  | int       | User ID of event creator           |
| capacity      | int       | Max attendees (optional)           |

---

## Environment Variables

| Variable      | Example Value                | Purpose                        |
|-------------- |-----------------------------|--------------------------------|
| APP_ENV       | development / production     | App environment                |
| PORT          | 8080                        | HTTP server port               |
| LOG_LEVEL     | debug / info / warn / error | Logging verbosity              |
| JWT_SECRET    | some-long-random-secret     | JWT signing secret             |
| DATABASE_URL  | (if using a DB)             | DB connection string           |

---

## Local Build & Run

1. **Get dependencies and build:**
     ```sh
     go mod tidy
     go build ./...
     ```
2. **Run the app locally (example):**
     ```sh
     export PORT=8080
     export JWT_SECRET=devsecret
     go run ./cmd/server
     ```
     - If the project doesn't have `cmd/server`, use `go run main.go` from the repository root or the package that contains `main()`.

---

## Testing

- **Unit tests:**
    ```sh
    go test ./... -v
    ```
- Add table-driven tests for handlers and small helper packages
- Mock external dependencies (DB, time, random) using interfaces

**Quick smoke test using curl (after app is running):**
```sh
curl -i http://localhost:8080/events
```

---

## Observability & Logging

- Use structured logging (`logrus`, `zerolog`, or standard library with JSON output)
- Keep logs consistent across packages
- Add request-level logging middleware to record method, path, status, and latency

---

## Authentication & Security

- Use JWT for stateless sessions. Keep `JWT_SECRET` out of source control and use environment-based configuration in CI/CD
- Validate user input rigorously and return clear error messages
- For production, ensure TLS termination at the load balancer or proxy

---

## Persistence & Data

- For a demo or learning project, a simple file-backed store or SQLite is acceptable
- For production, prefer PostgreSQL (managed) and use migrations (`golang-migrate` or similar)

---

## Dockerfile Example

```dockerfile
FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /bin/events-app ./cmd/server

FROM alpine:3.18
RUN addgroup -S app && adduser -S app -G app
COPY --from=builder /bin/events-app /bin/events-app
USER app
EXPOSE 8080
ENTRYPOINT ["/bin/events-app"]
```

---

## CI/CD Recommendations

- Add a pipeline stage to run `go vet`, `golangci-lint` (or `staticcheck`), `go test ./...`, and `go build`
- If publishing Docker images, build and push artifacts from CI with semantic tags

---

## Developer Guidance

- Keep functions small and focused. Prefer clear names over clever tricks
- Add comments to explain non-obvious decisions, but prefer clear code to comments when possible
- Add unit tests for bugs you fix: tests prevent regressions and are the best documentation for expected behavior
- When unsure about an API design, sketch the request/response shapes and error codes in a short PR description

---

## Edge Cases & Considerations

- **Timezones and timestamps:** Store times in UTC, and accept/return RFC3339 strings
- **Concurrency:** Protect shared in-memory resources with mutexes or use concurrent-safe data stores
- **Large payloads:** Enforce request size limits and validate payloads early
