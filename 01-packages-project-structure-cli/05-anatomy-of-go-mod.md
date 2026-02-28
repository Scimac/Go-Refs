## go.mod File

- A `go.mod` file defines:
	1. **Module path** (import path base).
	2. **Go version** the module is built with.
	3. **Dependencies** (with versions).
	4. **Replace/Exclude rules** for dependencies.

```go
module example.com/myapp

go 1.25
toolchain go1.25.1

require (
    github.com/gin-gonic/gin v1.9.1
    github.com/stretchr/testify v1.9.0
    golang.org/x/sys v0.14.0 // indirect
)

replace github.com/gin-gonic/gin => ../local-gin

exclude github.com/bad/dependency v1.2.3

retract v1.0.0
```
### Module Declaration

- `module` keyword → defines the **import path prefix** for your project.
- All local packages in your project are imported using this path. 
	`import "example.com/myapp/pkg/utils"`

```go
module example.com/myapp
```

### Go Version

```go
go 1.25
```

- Declares the minimum Go toolchain version needed.
- Ensures consistent builds (if someone uses an older Go version, they may see warnings or errors).

### Dependencies (`require`)

```go
require (
    github.com/gin-gonic/gin v1.9.1
    github.com/stretchr/testify v1.9.0
)
```

- Lists external dependencies + their versions.
- Version follows **Semantic Versioning** (`vMAJOR.MINOR.PATCH`).
- Direct dependencies (your code imports them) and indirect dependencies (imported by other deps) can appear here.

### Indirect Dependencies

```go
require (
    golang.org/x/sys v0.14.0 // indirect
)
```

- `// indirect` means your code **doesn’t import it directly**.
- Instead, another dependency (e.g., `gin`) requires it.
- Usually cleaned up with `go mod tidy`.

### Replace Directive

```go
replace github.com/old/module => ../local/module
```

- Tells Go to use a **different module path** or local copy.
- Common for:    
    - Using a **local version** during development.
    - Pointing to a **forked dependency**.

### Exclude Directive

```go
exclude github.com/bad/dependency v1.2.3
```

- Prevents a specific version of a dependency from being used.
- Useful if a version is buggy or has security issues.

### Toolchain Directive (Go 1.21+)

```go
toolchain go1.22.0
```

- Optional directive specifying the **recommended Go toolchain version**.
- Ensures developers use the same Go toolchain, even if they have an older one installed.

### Retract Directive

```go
retract v1.0.0
```

- Used in **published modules** to mark versions as invalid.
- For example, if `v1.0.0` was released with a security bug, you can retract it.

## go.sum File (Companion to go.mod)

- Stores **checksums** of dependencies (for security & reproducibility).
- Ensures builds are **verifiable** (same dependency versions everywhere).
- You usually don’t edit `go.sum` manually — it’s maintained by Go tooling.

## Runtime Ops

When you **build** or **run** a Go project:
1. Go looks at your **`go.mod`** file → to see what dependencies you declared.
2. Go checks **`go.sum`** → to verify the exact versions & integrity (checksums).
3. Dependencies are downloaded (if needed) into the **module cache**: (default:`~/go/pkg/mod`)

- Your code + dependencies are compiled together.

```shell
 ┌────────────────────────────┐
 │        Your Project        │
 │   go.mod   |   go.sum      │
 │   (deps)   | (checksums)   │
 └───────┬─────────┬──────────┘
         │         │
         ▼         ▼
 ┌───────────────┐   ┌────────────────┐
 │   Dependency  │   │ go.sum checks  │
 │   Versions    │   │   Integrity    │
 │ (require...)  │   │ (SHA256 hash)  │
 └───────────────┘   └────────────────┘
         │
         ▼
 ┌──────────────────────────────┐
 │  Module Cache (~/go/pkg/mod) │
 │  ├── github.com/gin-gonic/...│
 │  ├── github.com/stretchr/... │
 │  └── golang.org/x/...        │
 └──────────────────────────────┘
         │
         ▼
 ┌─────────────────────────────────┐
 │        Build System             │
 │   go build / go run / go test   │
 └─────────────────────────────────┘
```
