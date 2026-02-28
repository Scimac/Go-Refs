## Arrays

- An **array** is a fixed-size collection of elements of the **same type**.

```go
var arr [5]int

// output
// fmt.Println(arr) // [0 0 0 0 0]
```

- Size is **part of the type**
- `[5]int` ≠ `[6]int`
- Zero-valued by default

### Initialization

```go
arr1 := [3]int{1, 2, 3} // fixed length
arr2 := [...]int{4, 5, 6} // compiler infers length
```

### Characteristics of Arrays

|Property|Arrays|
|---|---|
|Size|Fixed|
|Mutability|Mutable|
|Passed to functions|By value (copied)|
|Usage|Rare in application code|
### Array Copy Semantics

- As arrays are value type DS in GO, They are copied when they are passed to the functions.

```go
func modify(a [3]int) {
    a[0] = 100
}

func main() {
    arr := [3]int{1, 2, 3}
    modify(arr)
    fmt.Println(arr) // [1 2 3]
}
```

### Use Cases: Arrays

- Low-level code
- Memory-critical systems
- Fixed-size buffers
- Rarely used in day-to-day Go

-  **In real Go code, arrays exist mainly to back slices.**

## Slices

- A **slice** is a **dynamic, flexible view** over an underlying array.

```go
s := []int{1, 2, 3}

// output
// [1 2 3]

sl := []int{}

// []
```

### Slice Internals
#important 

- Slice as a DS is a `Reference Type`. A slice is a **descriptor** with three parts:

```go
// struc tof slice in go
type slice struct {
    ptr *T  // pointer to underlying array
    len int
    cap int
}
```

- Slices **do not store data themselves** — they reference an array.

```go
fmt.Println(len(s)) // number of elements
fmt.Println(cap(s)) // capacity of underlying array
```

### Intialization

#### Slice Literal

```go
s := []int{1, 2, 3}
```

- Makes a slice based on Array of length 3.
#### Using make function

```go
s := make([]int, 5)       // len=5, cap=5
s := make([]int, 5, 10)   // len=5, cap=10
```

- By setting larger capacity, the base array referenced is initialized with larger length.
- This helps as `append` method does not need to create a new copy of array after appending.
- Saves operation cost, Better approach if length of the slice is known beforehand.
#### Using Array

- Technically, a slice should be based on an array. Therefore, it can be "sliced" out of an Array.

```go
arr := [5]int{1, 2, 3, 4, 5}
s := arr[1:4] // [2 3 4]
```

### Slice Operations

```go
// indexing
s[0] = 10

// slicing
s[a:b]   // from index a to b-1
s[:b]    // from 0 to b-1
s[a:]    // from a to end
s[:]     // full slice

// appending
s := []int{1, 2}
s = append(s, 3, 4)
```

- `append` **may or may not** allocate a new array.  #InterviewQs 
- Referring [[02-Arrays-slices-in-go#Intialization#Using make function| Slice using make]], of capacity of the base array is larger, then new array is not allocated by append.

### Capacity Growth
#InterviewQs

- Typical growth pattern:
	- Small slices → grow ~2x
	- Larger slices → grow ~1.25x (implementation detail)

```go
s := []int{}
for i := 0; i < 10; i++ {
    s = append(s, i)
    fmt.Println(len(s), cap(s))
}
```

### Slice Copy Semantics
#important #InterviewQs 

```go
a := []int{1, 2, 3}
b := a

b[0] = 100
fmt.Println(a) // [100 2 3]
```

- Slice Assignment Shares Memory. Both `a` and `b` point to the **same array**.
#### Making a Copy Properly

```go
// using copy function
src := []int{1, 2, 3}
dst := make([]int, len(src))
copy(dst, src)

// using append and ... operator
dst := append([]int{}, src...)
```

### Slices as Function Arguments

- Slices are **passed by value**, but the value contains a **pointer**. Modifying elements affects the caller.

```go
func modify(s []int) {
    s[0] = 100
}

func main() {
    a := []int{1, 2, 3}
    modify(a)
    fmt.Println(a) // [100 2 3]
}
```

- But Changing Length Does NOT Affect Caller. 

```go
func add(s []int) {
    s = append(s, 10)
}

func main() {
    a := []int{1, 2}
    add(a)
    fmt.Println(a) // [1 2]
}
```

- Re-slicing or appending **does not update caller’s slice header**.

### Nil Slice vs Empty Slice
#InterviewQs 

```go
var s1 []int        // nil slice
s2 := []int{}       // empty slice
```

|Property|nil slice|empty slice|
|---|---|---|
|Value|nil|non-nil|
|len|0|0|
|cap|0|0|
|JSON|null|[]|
- Prefer **nil slices** unless you need allocation.

### Removing Elements from Slice

#### Remove by index

```go
i := 2
s = append(s[:i], s[i+1:]...)
```

- Order preserved, O(n)

#### Fast Remove

```go
i := 2
s[i] = s[len(s)-1]
s = s[:len(s)-1]
```

- Order not preserved, O(1)

### Multidimensional Slices

```go
matrix := make([][]int, 3)
for i := range matrix {
    matrix[i] = make([]int, 4)
}
```

- Rows are **independent slices** (not contiguous like C).

## TODOs

- Implement for practice

| DS Problem     | Use                  |
| -------------- | -------------------- |
| Stack          | slice + append       |
| Queue          | slice or ring buffer |
| Two pointers   | slice                |
| Sliding window | slice                |
| Sorting        | slice + `sort`       |
| DP arrays      | slice                |

## Mental Model to Remember

> **Array = actual data**  
> **Slice = window into array**

