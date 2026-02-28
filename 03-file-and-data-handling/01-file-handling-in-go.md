- Go provides the `os`, `io`, `bufio`, and `ioutil` (deprecated in Go 1.16, replaced by `os`/`io`) packages for working with files.

### Opening a File

- `os.Open` → opens in **read-only mode**.
- Always `defer file.Close()` to release resources.

```go
file, err := os.Open("example.txt") // read-only
if err != nil {
    log.Fatal(err)
}
defer file.Close() // ! important
```

### Creating a File

- `os.Create` → creates file with **write-only + truncate**.
- Overwrites existing file.

```go
file, err := os.Create("newfile.txt") // truncates if exists
if err != nil {
    log.Fatal(err)
}
defer file.Close()
```

### Open File with Flags

```go
file, err := os.OpenFile("data.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
if err != nil {
    log.Fatal(err)
}
defer file.Close()
```
- Use `os.OpenFile` with flags for flexible control.

- Common flags:
    - `os.O_RDONLY` → read only    
    - `os.O_WRONLY` → write only    
    - `os.O_RDWR` → read & write        
    - `os.O_CREATE` → create if not exists
    - `os.O_APPEND` → append to file
    - `os.O_TRUNC` → truncate file
- `0644` → file permission (`rw-r--r--`).

### Reading from a File

```go
// Reading all the bytes
data, err := os.ReadFile("example.txt")
if err != nil {
    log.Fatal(err)
}
fmt.Println(string(data))
```

```go
// Using a Buffer
buffer := make([]byte, 100)
n, err := file.Read(buffer)
if err != nil && err != io.EOF {
    log.Fatal(err)
}
fmt.Println(string(buffer[:n]))
```
- Handle `io.EOF` when reading until end of file.

```go
// Line-by-Line (Buffered)
scanner := bufio.NewScanner(file)
for scanner.Scan() {
    fmt.Println(scanner.Text())
}
if err := scanner.Err(); err != nil {
    log.Fatal(err)
}
```
- Use `bufio.Scanner` for line-by-line reading (efficient).
### Writing to a File

```go
// Write String
file, err := os.Create("write.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

_, err = file.WriteString("Hello, Go!\n")
if err != nil {
    log.Fatal(err)
}
```

```go
// Append Mode
file, err := os.OpenFile("append.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
if err != nil {
    log.Fatal(err)
}
defer file.Close()

_, err = file.WriteString("New line\n")
if err != nil {
    log.Fatal(err)
}
```

### Copying Files

```go
src, _ := os.Open("source.txt")
defer src.Close()

dst, _ := os.Create("target.txt")
defer dst.Close()

_, err := io.Copy(dst, src)
if err != nil {
    log.Fatal(err)
}
```
- Use `io.Copy` for file duplication.
### Deleting / Renaming Files

```go
err := os.Remove("old.txt")         // delete file
err = os.Rename("old.txt", "new.txt") // rename file
```

### Checking File Info

```go
info, err := os.Stat("example.txt")
if err != nil {
    log.Fatal(err)
}
fmt.Println("Name:", info.Name())
fmt.Println("Size:", info.Size())
fmt.Println("IsDir:", info.IsDir())
```

### Reading & Writing JSON

```go
// writing json
data := map[string]string{"name": "Go", "type": "lang"}
file, _ := os.Create("data.json")
defer file.Close()
json.NewEncoder(file).Encode(data)
```

```go
// Reading JSON
file, _ := os.Open("data.json")
defer file.Close()
var result map[string]string
json.NewDecoder(file).Decode(&result)
fmt.Println(result)
```

### More Read

#### File Permissions in Go (`os.OpenFile`, `os.WriteFile`, etc.)

##### ## Structure of Permission Values

- File permission is represented as a **3-digit octal number** (sometimes 4-digit with a leading `0`).
- Each digit = **owner**, **group**, and **others** permissions.
- Digits are octal (base 8):
    
    - `4` = Read (`r`)
    - `2` = Write (`w`)    
    - `1` = Execute (`x`)    
    - Combine them with `+`.

##### Common Permission Examples
|Mode|Symbolic|Meaning|
|---|---|---|
|`0600`|rw-------|Only owner can read/write (secrets, keys).|
|`0644`|rw-r--r--|Owner read/write, others read-only (safe default for text).|
|`0666`|rw-rw-rw-|All users read/write (not secure).|
|`0755`|rwxr-xr-x|Owner can read/write/execute, others read/execute (executables).|
|`0777`|rwxrwxrwx|Everyone full access (avoid in production).|
- The **leading `0`** tells Go that the number is in **octal**.

##### Example

```go
package main

import (
	"log"
	"os"
)

func main() {
	// Create a file with 0644 permission (rw-r--r--)
	file, err := os.OpenFile("example.txt",
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
		0644,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString("Hello, permissions in Go!\n")
	if err != nil {
		log.Fatal(err)
	}
}
```

- File will be created with:
	- Owner → read/write (`rw-`)
	- Group → read-only (`r--`)
	- Others → read-only (`r--`)

##### Checking File Permissions

```cli
ls -l example.txt

// output
// -rw-r--r--  1 user group  25 Sep 20 23:59 example.txt
```