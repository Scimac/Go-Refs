- Since **Go 1.11**, Go introduced `modules` as the official dependency management system (replacing `GOPATH`).
- Modules are how you manage external libraries and versioning.

### Basic Commands

```bash
go mod init example.com/myapp   # Init new module with name example.com/myapp
go mod tidy                     # Add/remove dependencies automatically
go get github.com/gin-gonic/gin # Fetch a new dependency
go mod download                 # Download dependencies
go mod verify                   # Verify module dependencies
go list -m all                  # List all modules in use
```

### Project Structure

```bash
# A common convention (inspired by community best practices)
# Not enforced by GO, just a guideline
myapp/
│── go.mod            # Module definition & dependencies
│── go.sum            # Dependency checksum file
│── cmd/              # Application entry points (main packages)
│   └── myapp/
│       └── main.go
│── internal/         # Private packages (only usable inside this module)
│   └── service/
│       └── service.go
│── pkg/              # Publicly reusable packages
│   └── utils/
│       └── strings.go
│── api/              # API definitions (REST, gRPC, OpenAPI)
│── configs/          # Config files
│── scripts/          # DevOps/scripts
│── build/            # Build outputs, Dockerfiles, CI
│── test/             # Integration/e2e tests
```

- Use **modules** for dependency management (`go mod init`, `go mod tidy`). 
- Use **packages** to organize code (main vs library packages).
- Follow **project structures** (`cmd/`, `internal/`, `pkg/`) for clean architecture.
- Internal packages = private, pkg = public.