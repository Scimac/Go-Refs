#### Example: HTTP Package

- HTTP is a package in `net/http` package.
- `Get` is the function of `type Response` to fetch data from a particular URL. [doc](https://pkg.go.dev/net/http#Get)
```go
func Get(url string) (resp *Response, err error)
```

- Now the Response returned is not straight-forward JSON, as expected, but in a complex structure as mentioned in docs [Response](https://pkg.go.dev/net/http#Response)
- It is stored in `Body io.ReadCloser` key of the struct returned. Rest keys are KPIs of various other factors of the server/client viz. statue, protocol, header, contentLength, etc.

- The ReadCloser type in not a `concrete` type but a `interface` type. ReadCloser is the interface that groups the basic Read and Close methods.
```go
type ReadCloser interface {
	Reader
	Closer
}

type Reader interface {
	Read(p []byte) (n int, err error)
}
type Closer interface {
	Close() error
}
```

#### Reader Interface

- `Reader` interface defined a receiver function `Read` - that takes a `byte slice` and populates the response in the same slice. Also returns, 
	- `int` - number of bytes returned
	- `error` - errors returned if any.
-  When the end of the data is reached, `Read` returns `io.EOF` (Library constant). 

- The `Reader` interface in Go provides a standard way to read data from different sources.
- It allows you to use the same code to read from files, network connections, HTTP responses, in-memory buffers, and more.

**Common use cases:**

|Use Case|Example Type|Description|
|---|---|---|
|Reading from a file|`*os.File`|Read data from files on disk.|
|Reading HTTP responses|[resp.Body](vscode-file://vscode-app/Applications/Visual%20Studio%20Code.app/Contents/Resources/app/out/vs/code/electron-browser/workbench/workbench.html) ([io.Reader](vscode-file://vscode-app/Applications/Visual%20Studio%20Code.app/Contents/Resources/app/out/vs/code/electron-browser/workbench/workbench.html))|Read data from HTTP responses (as in your code).|
|Reading from memory|`*bytes.Buffer`|Read data from an in-memory buffer.|
|Reading from network|`net.Conn`|Read data from network sockets/connections.|
|Chaining readers|[io.TeeReader](vscode-file://vscode-app/Applications/Visual%20Studio%20Code.app/Contents/Resources/app/out/vs/code/electron-browser/workbench/workbench.html), [io.LimitReader](vscode-file://vscode-app/Applications/Visual%20Studio%20Code.app/Contents/Resources/app/out/vs/code/electron-browser/workbench/workbench.html)|Compose or modify streams (e.g., limit, duplicate, etc.).|

**Benefit:**  
- You can write functions that accept an [io.Reader](vscode-file://vscode-app/Applications/Visual%20Studio%20Code.app/Contents/Resources/app/out/vs/code/electron-browser/workbench/workbench.html) and work with any data source that implements this interface, making your code flexible and reusable.

**Reader in action**

```go
func bareReadWriteImplementation(resp *http.Response) {
	// read from response body
	// write to os.Stdout
	// this is a very basic implementation of io.Reader and io.Writer
	
	// Read from resp.Body takes in a byte slice on which it writes the data
	
	// it returns the number of bytes read and an error if any
	// Normally we use Writer Functions, instead of defining meaningless length of slice
	buf := make([]byte, 55000)
	for {
		n, err := resp.Body.Read(buf)
		if n > 0 {
			os.Stdout.Write(buf[:n])
		}
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error reading response body:", err)
			}
			fmt.Printf("the value of n in the last read is: %d\n", n)
			break
		}
	}
```

#### Write Interface

- The Write method writes the contents of p (a byte slice) to the underlying data stream.
- It returns the number of bytes written (n) and an error (err), if any.

```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

- We can implement a type that has Write function or use a already implemented one, like in `os.Stdout` type.

#### io.Copy function

- Instead of working through all this boilerplate code every time, [`io.copy`](https://pkg.go.dev/io#Copy)
- It accepts two types - a type with Writer function and a type with Reader function.
	```go
	func Copy(dst Writer, src Reader) (written int64, err error)
	```

- `io.Copy(dst, src)` reads from the i`o.Reader` (src) and writes to the `io.Writer`(dst) in a loop until all data is transferred (EOF).
- This abstracts away the manual buffer management and loop you wrote in `bareReadWriteImplementation`.

```go
func IoCopyExample(resp *http.Response) {
	// io.Copy is a more efficient way to copy from one to another
	// Output of the response is handled in the writer function logic of os.Stdout
	// Input of the response is handled in the reader function logic of resp.Body
	res, err := io.Copy(os.Stdout, resp.Body)

	if err != nil {
		fmt.Println("Error copying response body:", err)
	} else {
		fmt.Printf("\nNumber of bytes copied: %d\n", res)
	}
}
```

```go
+----------------+         +----------------+         +------------------+
|  Data Source   |         |   io.Reader    |         |   io.Writer      |
| (e.g. HTTP     | ----->  | (e.g. resp.Body| ----->  | (e.g. os.Stdout) |
|  response)     |         |   or file)     |         |   or file)       |
+----------------+         +----------------+         +------------------+
                                 |                             ^
                                 |                             |
                                 +-----------+-----------------+
                                             |
                                      io.Copy loop:
                                 +----------------------+
                                 | 1. Read from Reader  |
                                 | 2. Write to Writer   |
                                 | 3. Repeat until EOF  |
                                 +----------------------+
```

#### Custom Writer Function (Interface Pitfall)

- As mentioned before, we can pass our own type with Writer function.

```go
type logWriter struct{}

func (logWriter) Write(bs []byte) (int, error) {
	// custom writer to log the number of bytes written
	fmt.Println("This is a custom log writer")
	fmt.Println(string(bs))
	fmt.Printf("Just wrote this many bytes: %d\n", len(bs))
	return len(bs), nil
}

func CustomWriterExample(resp *http.Response) {
	lw := logWriter{}
	n, err := io.Copy(lw, resp.Body)
	if err != nil {
		fmt.Println("Error copying response body to custom writer:", err)
	} else {
		fmt.Printf("Number of bytes copied to custom writer: %d\n", n)
	}
}
```

- But we need to careful, because wrong logic can't be caught using just interfaces.

```go
func (logWriter) Write(bs []byte) (int, error) {
// some JUNK logic, wrong response
	return 1, nil
}
```

- So, be careful around the Interfaces. 