### Create Project Folder

- Create the project **anywhere** on your filesystem.
- No need for `$GOPATH/src`.

```bash
mkdir myapp
cd myapp
```

### Initialise Go Module

- Create `go.mod` file

```bash
go mod init example.com/myapp
```

- Module path (`example.com/myapp`) is used for importing packages within the project.

```go
// initial go.md
module example.com/myapp // module path

go 1.25
```

### Create Project Structure

```bash
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
│       └── greet.go
```

- `cmd/myapp` → executable entry point (`main` package)
- `internal/` → private packages for business logic
- `pkg/` → reusable public packages

### Add a Utility Package

```go
// pkg/utils/greet.go // file name can be anything
package utils // package name same as folder name

func Greet(name string) string {
    return "Hello, " + name + "!"
}
```

### Create `main.go`

```go
// cmd/myapp/main.go
package main

import (
    "fmt"
    "example.com/myapp/pkg/utils"
)

func main() {
    fmt.Println(utils.Greet("Makarand"))
}
```

### Run the Project

```bash
go run ./cmd/myapp
```

```bash
Hello, Makarand!
```

```bash
# can also build a binary
go build -o myapp ./cmd/myapp
./myapp
```

### Adding External Dependencies

```shell
# using Gin framework
go get github.com/gin-gonic/gin@v1.9.1
```
- Updates `go.mod` and creates/updates `go.sum`

```go
import "github.com/gin-gonic/gin"
```

### Tidying Dependencies

- Cleans up unused dependencies
- Ensures reproducible builds

```bash
go mod tidy
```

### best practices

- Keep **one `main()` per executable** in `cmd/`
- Use `internal/` for private code (cannot be imported outside module)
- Use `pkg/` for reusable libraries
- Always run `go mod tidy` after adding/removing dependencies
- Version control your module (`go.mod` + `go.sum`)

