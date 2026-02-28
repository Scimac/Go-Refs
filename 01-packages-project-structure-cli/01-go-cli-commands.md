- The `go` command is the main entry point to work with Go.
## Commands

### Help commands

- `go help` / `go help {command}`
	- Shows general usage and available subcommands.

### Run commands

- `go run` / `go run main.go`
	- Compiles and executes `main.go` directly (good for quick testing).
	- For multiple files - `go run file1.go file2.go`
- `go run .`
	- Runs the entire current directory (with a `main` package).
- `go build` / `go build -o app`
	- Compiles the current package into a binary (executable).

### Config / Stats Commands

- `go env`
	- Displays Go environment variables (`GOPATH`, `GOROOT`, etc).
- `go version`
	- Prints Go version installed.
- `go list ./...`
	- Lists all packages in your project.

### Dependency management commands

- `go mod init myproject`
	- Initializes a new module with name `myproject`.
- `go install`
	- Compiles and installs the binary into `$GOPATH/bin` or module bin path.
	- Useful when you want to reuse the binary globally.
- `go mod tidy`
	- Adds missing and removes unused dependencies in `go.mod`.
- `go get` / `go get github.com/gorilla/mux`
	- Downloads and installs a dependency.
- `go list -m all`
	- Lists all modules used in the current project.

### Testing commands

- `go test
	- Runs tests in the current package (files ending with `_test.go`).
- `go test ./...`
	- Runs tests in all sub-packages.
- `go test -v`
	- Verbose output (shows which tests ran).
- `go test -bench=.`
	- Runs benchmarks.

### Linter Commands

- `go fmt` / `go fmt ./...`
	- Formats code according to Go style.
- `go vet ./...`
	- Reports suspicious constructs (static analysis).

### Documents Commands

- `go doc fmt.Println`
	- Shows documentation for a function/package.
- `go doc -all net/http`
	- Shows all docs for `net/http`.

### Cache Commands

- `go clean -cache -modcache -i -r`
	- Cleans build and module cache.
- `go fmt ./... && go vet ./...`
	- Format + Vet before commit.
