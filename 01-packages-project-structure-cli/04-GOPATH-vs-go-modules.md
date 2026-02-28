### GOPATH

- All Go code had to live under a single workspace defined by `GOPATH`.
- Default `GOPATH` was `~/go` (on Linux/Mac).
- Source code → `$GOPATH/src`
- Binaries → `$GOPATH/bin`
- Packages → `$GOPATH/pkg`

```bash
# Directory layout

$GOPATH/
├── src/
│   └── github.com/user/project/
│       ├── main.go
│       └── utils/utils.go
├── bin/
└── pkg/
```

#### Issues

- **Global workspace**: only one `$GOPATH`. Hard to manage multiple projects.
- **Versioning problem**: couldn’t specify versions of dependencies easily.
- **Dependency conflicts**: upgrading a dependency broke all projects.

### Go Modules (Modern Workflow)

- Each project is **self-contained** with its own `go.mod` file.
- You can work **anywhere on disk** (no need for `$GOPATH/src`).
- Dependencies are downloaded into a global cache (`$GOPATH/pkg/mod`).
- Allows **versioning** and reproducible builds.

```bash
myapp/
├── go.mod
├── go.sum
├── cmd/myapp/main.go
└── pkg/utils/greet.go
```

#### Advantages

- No need for `$GOPATH` restrictions.
- Explicit version control of dependencies.
- Works well with **semantic versioning**.
- Enables reproducible builds.

- Even though modules are default now, Go still respects `$GOPATH` for:
	- Binary installs (`go install`)
	- Module cache (`$GOPATH/pkg/mod`)

