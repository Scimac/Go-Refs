- A **map** is a built-in data structure that stores **key → value** pairs.

```go
map[KeyType]ValueType

m := map[string]int{
    "apple":  3,
    "banana": 5,
}
```

- Implemented as a **hash table**
- Unordered
- Reference type
- Extremely fast for lookups - O(1)
- Insert Operation - O(1)
- Delete Operation - O(1)
- Zero Value - `nil`

### Initialization

#### Using make function

```go
m := make(map[string]int)

m := make(map[string]int, 100)
```

- Capacity hint improves performance by reducing rehashing.
#### Using map Literal

```go
m := map[int]string{
    1: "one",
    2: "two",
}
```

### Map Internals
#important  #InterviewQs 

- A Go map variable is a pointer to an internal `hmap` structure in the runtime, which contains metadata about the map: 
	- **Count:** The number of elements currently in the map.
	- **Number of buckets:** Stored as a log base 2 of the actual number of buckets.
	- **Hash seed:** A random seed used by the hash function to ensure different maps have different hash behavior, which helps prevent certain types of attacks and ensures random iteration order.
	- **Pointers to buckets:** A pointer to the primary array of buckets and an optional pointer to an old bucket array during a resize operation.
- Buckets (`bmap` struct) - The primary storage unit of the map is an array of buckets. Each bucket is designed to hold a limited number of key-value pairs to handle collisions efficiently: 
	- **Key-Value Pairs:** Each bucket can store up to **8** key-value pairs.
	- **Top Hash Values:** An array of hash code snippets (specifically the top 8 bits) for each key in the bucket, used for faster lookups within the bucket without comparing the full key every time.
	- **Overflow Pointer:** If a bucket exceeds its capacity (e.g., more than 8 elements hash to the same bucket index), an overflow pointer links to an additional "overflow" bucket, creating a chain to handle collisions.

- Conceptually:
	- Hash function → bucket    
	- Buckets contain key-value pairs
	- Automatic resizing & rehashing    

- You **cannot**:
	- Control hash function
	- Control bucket count
	- Get capacity (no `cap(m)`)
- But you **can** give initial size hint:
```go
make(map[string]int, 1000)
```

### Zero Value: `nil` Map
#InterviewQs 

```go
// nil Map
var m map[string]int

fmt.Println(m) // nil, not {}
fmt.Println(m["x"]) // 0 (zero value)
m["x"] = 10         // panic

// making a empty value map
mp := map[string]string{}
```

- Always `make` a map before writing.

### Basic Operations

```go
// insert
m["go"] = 10

// read
v := m["go"]

// delete
delete(m, "go")
```

### Copying Maps

- Maps cannot be copied directly.

```go
b := a // does not copy a map

// Manual COpy
copyMap := make(map[string]int, len(a))
for k, v := range a {
    copyMap[k] = v
}
```

### Maps Are Reference Types

- Assignment Shares Memory
```go
a := map[string]int{"x": 1}
b := a

b["x"] = 100
fmt.Println(a["x"]) // 100

// Both `a` and `b` refer to the same map.
```

- Passing Maps to Functions
```go
func modify(m map[string]int) {
    m["x"] = 10
}

func main() {
    m := map[string]int{}
    modify(m)
    fmt.Println(m["x"]) // 10
}

// Changes are visible to caller.
```

### The “Comma OK” Idiom
#important 

- Used to distinguish **missing key vs zero value**.

```go
v, ok := m["java"]
if ok {
    fmt.Println("Found:", v)
} else {
    fmt.Println("Not found")
}
```

- Without `ok`, you **cannot tell the difference**.

```go
m["x"] = 0

v := m["y"]  // also returns 0
```

### Iterating in Maps
#important  

- **Iteration order in map is random.** Go deliberately randomizes map iteration to prevent bugs. Never rely on map order.

```go
for k, v := range m {
    fmt.Println(k, v)
}
```

#### Deleting While Iterating

```go
for k := range m {
    delete(m, k) // allowed and safe
}
```

### Map Keys: What Is Allowed?

-  Keys MUST be `comparable`
	- Allowed:
		- `int`, `string`, `bool`
		- `struct` (if all fields comparable)
		- pointers
		- arrays
	- Not allowed:
		- `slice`
		- `map`
		- `function`

### Maps with Struct Values

```go
type User struct {
    Name string
    Age  int
}

users := map[int]User{
    1: {"Alice", 30},
}
```

- Modifying struct field directly does NOT work:
```go
users[1].Age = 31 // compile error
```

- correct way
```go
// update the struct variable directly
u := users[1]
u.Age++
users[1] = u

// or pass a pointer
map[int]*User
```

### Common Map Patterns
#important 
#### Frequency Counter

```go
freq := map[string]int{}
for _, w := range words {
    freq[w]++
}
```

#### Set Implementation

```go
set := map[string]struct{}{}
set["go"] = struct{}{}

// struct used because - Zero memory allocation - Semantic clarity
```

#### Grouping

```go
groups := map[string][]int{}
groups["even"] = append(groups["even"], 2)
```

#### Graph Representation

```go
graph := map[int][]int{
    1: {2, 3},
    2: {4},
}
```

### Practic Problems

- Implement the following using maps - 
	- Two Sum
	- Anagrams
	- Longest substring
	- LRU Cache (map + list)
	- Graph traversal

