- In Go, **asynchronous programming = concurrency**, unlike callbacks or promises in JS.

> Go uses **goroutines + channels**, not `async/await`.

- Key idea:
	- Many tasks **in progress at the same time**
	- Managed by Go runtime, **not OS threads**

## Goroutines — Lightweight Concurrent Units

- A **goroutine** is a function executing concurrently with other goroutines.
```go
// Add go keyword in front of the function call to make it concurrent
	go myFunction()
```

### Example

```go
func hello() {
    time.Sleep(time.Second)
    fmt.Println("Hello from goroutine")
}

func main() {
    go hello()
}

// output:
//
```

- Goroutine makes function call async, but it doesn't hold the `main` function execution waiting for the result to come!!
- Program exits when `main` exits — goroutines don’t keep it alive.

### Goroutines vs Threads
#InterviewQs 

| Aspect        | Goroutine     | OS Thread |
| ------------- | ------------- | --------- |
| Creation cost | Very cheap    | Expensive |
| Stack size    | ~2 KB (grows) | ~1–8 MB   |
| Scheduling    | Go runtime    | OS        |
| Count         | Millions      | Thousands |
## Go Schedulers- G-M-P Scheduler Model

- Goroutines are **cheap, managed, multiplexed tasks**.
- Go uses **M:N scheduling**.
	- **G** → **Goroutine** - A lightweight function or task that is waiting to execute. It's the unit of concurrency in Go.
	- **M** → **OS Thread** - A real operating system thread managed by the OS kernel.
	- **P** → **Logical processor** - A logical context that holds a local run queue of goroutines and other scheduler state.
- Multiplexing M goroutines onto N operating system threads, using a P (Processor) as an execution context.
- Go runtime schedules goroutines across threads automatically.

```txt
Many G  → few M → CPU cores
```

## Channels

> Do not communicate by sharing memory; share memory by communicating.

- Channels are **typed conduits** for goroutines to talk safely.

```go
ch := make(chan int) // This channel will only accept int type
```

```go
func worker(ch chan int) {
    ch <- 42
}

func main() {
    ch := make(chan int)
    go worker(ch)
    value := <-ch
    fmt.Println(value)
}
```

- Synchronization + communication together. No locks needed.

### Channel Direction (Compile-time Safety)

```go
func send(ch chan<- int) { ch <- 1 }
func receive(ch <-chan int) { fmt.Println(<-ch) }
```

### Blocking Behaviour

- This is how Go achieves synchronization.

1. Send Blocks - Blocks until someone receives.
   ```go
	     ch <- x
	```
2. Receive Blocks - Blocks until someone sends.
   ```go
	   x := <-ch
	```

### Buffered vs Unbuffered Channels

- In Go, the primary difference between **buffered and unbuffered channels** lies in - 
	- their capacity to store values and 
	- their blocking behavior, which affects how goroutines synchronize.
- Unbuffered channels enforce **synchronous communication**, while buffered channels allow a degree of **asynchronous communication** by using a queue. 

#### Unbuffered Channel (default)

- An unbuffered channel has a capacity of zero and cannot store any data. Communication (send and receive operations) only succeeds when both a sender and a receiver are ready at the same time.

- Sender waits for receiver
- Strong synchronization

```go
ch := make(chan int)
```

#### Buffered Channel

- A buffered channel has a fixed capacity greater than zero to store a limited number of values. 
- It acts as a queue, allowing a sender to continue sending as long as the buffer is not full, without waiting for an immediate receiver.

- Sender blocks only when buffer is full
- Receiver blocks when buffer empty

```go
ch := make(chan int, 3)
```

## `select` Statement — Async Control Center

- `select` lets you wait on **multiple channel operations**.

```go
select {
case v := <-ch1:
    fmt.Println(v)
case ch2 <- 10:
    fmt.Println("sent")
default:
    fmt.Println("non-blocking")
}
```

### Blocking Select

- Blocks until one case is ready.
```go
select {
	case <-ch1:
	case <-ch2:
}
```

#### Timeout Pattern
#InterviewQs 

```go
select {
case v := <-ch:
    fmt.Println(v)
case <-time.After(time.Second):
    fmt.Println("timeout")
}
```

## Goroutine Lifecycle & Leaks

#### Goroutine Leak (COMMON BUG)

```go
func worker(ch chan int) {
    for {
        <-ch
    }
}
```

- If `ch` is never closed → goroutine blocks forever.
- **Leaked goroutines consume memory**
- Solution: Use `close` or `context`

## Closing Channels

- Only sender should close 
- Receivers detect closure

```go
close(ch)
```

- An attempt to close a `nil` channel, an already closed channel, or a channel from the receiver side will cause a runtime `panic`.

- Use `ok` return param to check if the 
```go
v, ok := <-ch
if !ok {
    fmt.Println("channel closed")
}
```

- If `ok` is `true`, a value was successfully received. 
- If `ok` is `false`, the channel is closed and no more values are available (and `val` will be the zero value). 

### Ranging over a channel in GO

- This is particularly useful for terminating a `for range` loop over a channel, as the loop automatically stops when the channel is closed and drained.

```go
for v := range ch {
    fmt.Println(v)
}
```

- This stops automatically when channel closes. As Channels are also Auto GC'ed.

