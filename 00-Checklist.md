## 1. Basics of Go

- ✅ CLI (`go run`, `go build`, `go fmt`, `go mod`, etc.)
- ✅ Packages (`main`, imports, stdlib overview)
- ✅ Variables & Constants (explicit/implicit, zero values, conversions)
- ✅ `fmt` & Formatted Strings
---
## 2. Core Go Language

- ✅ Functions (multiple returns, variadic, methods, defer, error handling)
- ✅ Conditionals & Loops (`if/else`, `switch`, `for`, labels, control flow)
- ✅ Data Structures: arrays, slices, maps, structs
- ✅ Pointers & Memory Management
- ✅ Interfaces & Type Embedding
- ✅ Error Handling patterns (`errors`, `fmt.Errorf`, custom errors)

- **Pointers and Memory Management** (references, dereferencing, best practices).
- **Structs and Methods** (object-like behavior).
- **Interfaces** (Go’s way of abstraction).
- **Concurrency** (goroutines, channels).
---
## 3. File & Data Handling

- ✅ File I/O (`os`, `io`, `bufio`)
- ✅ JSON encoding/decoding
- ➡️ CSV✅, XML, YAML parsing
- ➡️ Environment variables (`os.Getenv`)
---
## 4. Concurrency & Parallelism

- ➡️ Goroutines (lightweight threads)
- ➡️ Channels (sync & async communication)
- ➡️ Select statements
- ➡️ Mutexes & WaitGroups (`sync` package)
- ➡️ Context package (timeouts, cancellations)
---
## 5. Modules & Project Structure

- ✅ `go mod init`, `go get`
- ✅ Dependency management
- ✅ Organizing packages in real projects
---
## 6. Networking & Web Development

- ➡️ HTTP servers (`net/http`)    
- ➡️ REST API basics
- ➡️ Middleware (logging, auth) 
- ➡️ JSON APIs with `encoding/json`   
- ➡️ Popular frameworks: **Gin**, **Echo**, **Fiber**   
---
## 7. Database Handling

- ➡️ SQL in Go (`database/sql`)
- ➡️ Using Postgres/MySQL drivers
- ➡️ ORM libraries (GORM, sqlx)
- ➡️ Transactions, migrations
---
## 8. Advanced Backend Topics

- ➡️ Testing (`testing` package, testify, mocks)
- ➡️ Error handling best practices
- ➡️ Logging (`log`, `zap`, `logrus`)
- ➡️ Configuration management (env, config files, Viper)
- ➡️ Graceful shutdown & signals (`os/signal`, `context`)
---
## 9. Deployment & Performance

- ➡️ Building & cross-compilation (`go build`)
- ➡️ Binaries & Dockerizing Go apps
- ➡️ Profiling & Benchmarking (`pprof`, `testing.B`)
- ➡️ Performance optimization (goroutines, memory use)
---
## 10. Cloud & Production-Ready Go

- ➡️ GRPC services (`google.golang.org/grpc`)
- ➡️ Message queues (Kafka, NATS, RabbitMQ clients)
- ➡️ Microservices patterns in Go
- ➡️ Kubernetes & Cloud deployments