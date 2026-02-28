## Go Packages

- A package is a collection of related Go source files in the same directory.
- Package name decides how code can be imported and reused.

- Packages in Go can be distributed across multiple files.
- Also, One project in Go can have multiple packages.

> In Go, First line of each file should define which package does the file belong to. 
### Main Package

- `go mod init name` - initializes a module for us in the current directory. 
- It uses main package as an entry point for building project and creating a `.exe` file.

- `package main` → special package with `main()` entry point (compiles into an executable).
- The `func main ()` is executed first in the main package, and cannot have more than one `main` function declaration in the `main` package.

- Modules / Project that don't need to generate `exe` file (i.e reusable packages), do not need a `main` function.

```go
// cmd/myapp/main.go
package main

import (
    "fmt"
    "example.com/myapp/pkg/utils"
)

func main() {
    fmt.Println(utils.Reverse("hello"))
}
```

```go
// utils/strings.go
package utils

// Capital first letter means the function is exported
// no explicit 'export' keyword in Go 
func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
```
### Types of packages 

- Executable packages
	- Declared as `package main`.
	- Must contain `func main()`.
	- Produces a binary when compiled. go run/build main.go
- Reusable packages
	- Declared as anything **other than `main`**.
	- Provides reusable code.
	- Cannot be executed directly.
	- Example: `fmt`, `net/http`, or your custom `utils`.

### Package Visibility

-  **Exported identifiers**: Start with an **uppercase letter**.
    - Accessible outside the package.
    - Example: `fmt.Println`.

- **Unexported identifiers**: Start with a **lowercase letter**.
	- Only usable within the package.

```go
// mymath/mymath.go
package mymath

func Add(a, b int) int {  // exported
    return a + b
}

func subtract(a, b int) int {  // unexported
    return a - b
}
```

### Important Guidelines

- Keep **only one `main()` per executable** (under `cmd/` usually).
- Separate **business logic** (e.g., in `internal/`) from **infrastructure code**.
- Avoid very large packages — split into smaller logical packages.
- Keep package names **short, lowercase, no underscores**.
- Internal vs External Packages
	- **`internal/`**:  
	    Packages under `internal/` can only be imported by code **within the same module**.  
	    Prevents unwanted usage.
	- **`pkg/`**:  
	    Publicly reusable code. Other projects can import it.
## Import statements

- Packages can be imported like - 
```go
import "fmt"
// OR
import (
    "fmt"
    "math"
    "strings"
)
```

```go
// Special Imports like - 

// Alias import
import m "math"
fmt.Println(m.Sqrt(16))

// Dot import (not recommended)
import . "fmt"
Println("Hello")  // no fmt. prefix

// Blank import - Used only for side-effects (like DB drivers).
import _ "github.com/lib/pq"  // registers PostgreSQL driver
```

## Important Standard Packages

Here are **must-know packages** for backend + general Go usage:
### General Utilities

- `fmt` → formatted I/O (`fmt.Println`, `fmt.Sprintf`)
- `os` → OS interaction (files, env vars, args)
- `io` & `io/ioutil` → input/output operations
- `log` → logging
- `flag` → command-line args
### Strings & Data

- `strings` → string manipulation
- `strconv` → conversions (`string ↔ int`)
- `math` → math functions
- `time` → working with time/dates
- `encoding/json` → JSON encode/decode
### Web & Networking

- `net/http` → building HTTP servers & clients
- `net/url` → URL parsing & handling
- `html/template` → HTML templating
### Concurrency & Context

- `sync` → locks, `waitgroups`
- `context` → cancellation, timeouts
- `runtime` → system/runtime info (`goroutines`, GC)
### Testing

- `testing` → unit testing
- `testing/benchmark` → performance testing

