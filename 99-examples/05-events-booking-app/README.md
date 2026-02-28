# Learnings: Building a Go Events Booking API (Beginner-Friendly Guide)

Welcome! This document is designed for Go beginners. It explains the key concepts, patterns, and lessons from building a RESTful Events Booking API using Go, Gin, SQLite, and modern best practices. Each section introduces a concept, explains why it matters, and shows how it’s used in this project.

---


## 1. Planning the API & Database Schema

**See also:**
- [`routes/routes.go`](routes/routes.go) — API endpoint registration
- [`models/events.go`](models/events.go), [`models/users.go`](models/users.go), [`models/registrations.go`](models/registrations.go) — Table/struct definitions

**Why plan?**
Before writing code, it’s important to decide what your app will do (the API endpoints) and what data you’ll store (the database schema). This helps you avoid confusion and rework later.

**API Design:**
- List the actions your users need (e.g., create event, register, login).
- For each action, define an HTTP method (GET, POST, etc.) and a path (e.g., `/events`).

**Database Schema:**
- Decide what tables you need (e.g., `users`, `events`, `registrations`).
- For each table, list the columns and their types (e.g., `name TEXT`, `date DATETIME`).

---


## 2. Gin Framework & Route Handling

**See also:**
- [`routes/routes.go`](routes/routes.go) — Route setup and handler registration
- [`routes/users.go`](routes/users.go), [`routes/events.go`](routes/events.go) — Handler implementations

**What is Gin?**
Gin is a popular web framework for Go. It makes it easy to define HTTP routes and handlers.

**Key Concepts:**
- **Router:** Maps HTTP requests to handler functions.
- **Handler:** A function that processes a request and returns a response.
- **Middleware:** Functions that run before/after handlers (e.g., for authentication).

**Example:**
```go
server := gin.Default()
server.GET("/events", getAllEvents) // GET /events handled by getAllEvents
```

---


## 3. Setting Up SQLite Database

**See also:**
- [`db/db.go`](db/db.go) — Database initialization and table creation
- [`models/events.go`](models/events.go), [`models/users.go`](models/users.go) — DB usage in models
## 3a. Preparing Queries vs. Direct Execution in Go SQL

**See also:**
- [`models/events.go`](models/events.go) — Examples of Prepare+Exec and direct Exec/Query

**Why SQLite?**
SQLite is a lightweight, file-based database. It’s great for learning and small projects.

**Key Steps:**
- **Initialize the DB:** Open a connection to the database file.
- **Connection Pooling:** Set max open/idle connections for performance.
- **Create Tables:** Use SQL to define your schema.

**Go SQL Package:**
- `DB.Exec()`: Run SQL commands that don’t return rows (e.g., CREATE TABLE).
- `DB.Prepare() + stmt.Exec()`: Prepare a statement for repeated use (good for performance if used multiple times).
- `DB.Query()`: Run SQL commands that return rows (e.g., SELECT).

**Example:**
```go
stmt, err := db.Prepare("INSERT INTO events (name, date) VALUES (?, ?)")
stmt.Exec("Go Meetup", "2026-02-15T10:00:00Z")
```

**Note:** In this project, we often close the statement after one use, so `Prepare()` is not strictly needed, but it’s good to know for larger apps.

## 3a. Preparing Queries vs. Direct Execution in Go SQL

**Direct Execution (`Exec`/`Query`):**
- Use `DB.Exec()` for commands that don’t return rows (e.g., INSERT, UPDATE, DELETE, CREATE TABLE).
- Use `DB.Query()` for commands that return rows (e.g., SELECT).
- These methods are simple and great for one-off queries.

**Example:**
```go
// Directly insert a new event
_, err := db.Exec("INSERT INTO events (name, date) VALUES (?, ?)", "Go Meetup", "2026-02-15T10:00:00Z")
```

**Preparing a Statement (`Prepare` + `Exec`):**
- `Prepare` creates a reusable SQL statement with placeholders (?).
- Useful if you’ll run the same query many times (e.g., in a loop).
- Can improve performance by parsing the SQL only once.

**Example:**
```go
stmt, err := db.Prepare("INSERT INTO events (name, date) VALUES (?, ?)")
defer stmt.Close()
for _, event := range events {
    _, err := stmt.Exec(event.Name, event.Date)
    // handle err
}
```

**Which should you use?**
- For a single query, use `Exec` or `Query` directly.
- For repeated queries (especially in a loop), use `Prepare` + `Exec` for better performance.
- In this project, we often use `Prepare` for demonstration, but for one-time inserts, direct `Exec` is fine.

**Tip:** Always close your prepared statements with `defer stmt.Close()` to avoid resource leaks.
---


## 4. Password Hashing with bcrypt

**See also:**
- [`utils/password.go`](utils/password.go) — Password hashing helpers (if present)
- [`models/users.go`](models/users.go) — User password handling

**Why hash passwords?**
Never store plain-text passwords! Hashing protects user data if your DB is leaked.

**bcrypt:**
- A secure hashing algorithm designed for passwords.
- In Go, use the `golang.org/x/crypto/bcrypt` package.

**Example:**
```go
hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
```

---


## 5. Working with JWT Tokens

**See also:**
- [`utils/jwt.go`](utils/jwt.go) — JWT creation and verification
- [`middlewares/auth.go`](middlewares/auth.go) — JWT middleware usage

**What is JWT?**
JWT (JSON Web Token) is a way to securely transmit information (like user ID) between parties as a JSON object.

**Why JWT?**
- Used for stateless authentication (no session storage needed).

**Key Concepts:**
- **Claims:** Data inside the token (e.g., user ID, expiration time).
- **exp:** The expiration time claim (when the token becomes invalid).

**Example:**
```go
token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    "user_id": user.Id,
    "exp": time.Now().Add(time.Hour * 24).Unix(),
})
```

---


## 6. Adding Middlewares in Gin

**See also:**
- [`middlewares/auth.go`](middlewares/auth.go) — Example middleware
- [`routes/routes.go`](routes/routes.go) — Applying middleware to route groups

**What is middleware?**
Middleware is code that runs before or after your main handler. It’s used for things like authentication, logging, or error handling.

**How to use in Gin:**
- You can add middleware globally or to specific route groups.

**Example:**
```go
authGroup := server.Group("/")
authGroup.Use(AuthMiddleware)
authGroup.POST("/events", createEvent)
```

**Multiple Handlers:**
- Gin lets you chain multiple handlers. They run left to right.

**Grouping Routes:**
- Use groups to apply common middleware to related routes.

---


## 7. Practical Tips & Gotchas

**See also:**
- [`models/events.go`](models/events.go), [`models/users.go`](models/users.go) — Value vs pointer receivers, error handling
- [`utils/jwt.go`](utils/jwt.go) — Type assertions in JWT parsing

- **Use clear, consistent naming for routes and handlers.**
- **Always check errors** when working with the database or external packages.
- **Understand value vs pointer receivers** in Go methods (pointer receivers let you modify the struct).
- **Type assertions:** When decoding JSON, numbers may come as float64, not int64—convert as needed.
- **Debugging:** Use print statements or a debugger (like Delve or VS Code’s Go extension) to step through code.

---


## 8. Summary Table: Key Go Concepts

**See also:**
- [`models/`](models/) — Structs, methods, error handling
- [`routes/`](routes/) — Handlers, middleware

| Concept         | What it is / Why it matters                | Example / Usage                |
|-----------------|--------------------------------------------|-------------------------------|
| Struct          | Custom data type (like a class)            | `type User struct { ... }`     |
| Method          | Function with a receiver (like a class fn) | `func (u *User) Save() error`  |
| Package         | Group of related code/files                | `package models`               |
| Interface       | Defines behavior, not data                 | `type Reader interface { ... }`|
| Error handling  | Idiomatic in Go: always check `err`        | `if err != nil { ... }`        |
| Pointer         | Reference to a value, allows mutation      | `func (e *Event) Save()`       |
| Middleware      | Code that wraps handlers                   | `server.Use(Logger)`           |
| JWT             | Auth token, stateless sessions             | `jwt.NewWithClaims(...)`       |
| bcrypt          | Password hashing                           | `bcrypt.GenerateFromPassword`  |

---


## 9. Next Steps for Learners

**See also:**
- [`api-test/`](api-test/) — HTTP client examples for testing your APIs

- Try adding a new API endpoint (e.g., update user profile).
- Experiment with error handling: what happens if the DB is down?
- Add logging to your handlers.
- Write unit tests for your models and handlers.
- Explore Go’s documentation and the Gin framework docs.

---

**You’re on your way to becoming a Go developer!**
Keep experimenting, break things, and ask questions. Every bug is a learning opportunity.

