### Value Type

- Copy happens on **assignment**
- Copy happens on **function call**
- Each variable owns its **own data**

#### Built-in Primitive Value Types

| Type                | Notes                |
| ------------------- | -------------------- |
| `int`, `int64`, etc | Numeric              |
| `float64`           | Numeric              |
| `bool`              | Boolean              |
| `string`            | Immutable value type |
| `array`             | Fixed-size value     |
| `struct`            | Composite value      |

#### Operations on Value type

```go
a := 10
b := a

b = 20
fmt.Println(a) // 10
fmt.Println(b) // 20
// Entirely new int created

arr1 := [3]int{1, 2, 3}
arr2 := arr1

arr2[0] = 100
fmt.Println(arr1) // [1 2 3]
fmt.Println(arr2) // [100 2 3]
// Entirely new array created here after copy

type User struct {
    Age int
}

u1 := User{Age: 30}
u2 := u1

u2.Age = 40
fmt.Println(u1.Age) // 30
fmt.Println(u2.Age) // 30
// Entirely new struct created here after copy
```

### Reference Type
- Assignment copies a **reference**
- Multiple variables can point to **same underlying data**
- Mutation affects all references

- These types **do not store data directly**, but reference underlying data.
#### Built-in Referenced Types in Go

| Type        | Why Reference                 |
| ----------- | ----------------------------- |
| `slice`     | Header points to array        |
| `map`       | Pointer to runtime hash table |
| `channel`   | Pointer to runtime queue      |
| `function`  | Pointer to code + closure     |
| `interface` | (type, value) pair            |
#### Operations on Referenced types

##### Slice Header (Conceptual)

```go
type slice struct {
    ptr *T
    len int
    cap int
}
```

```go
a := []int{1, 2, 3}
b := a

b[0] = 100
fmt.Println(a) // [100 2 3]
fmt.Println(b) // [100 2 3]
// slice a got changed too, as Both `a` and `b` point to the same underlying array.
// Copying a slice copies only this header.

m1 := map[string]int{"x": 1}
m2 := m1

m2["x"] = 100
fmt.Println(m1["x"]) // 100
// Same underlying hash table.

x := 10
p := &x

*p = 20
fmt.Println(x) // 20
// Pointer holds **address of value**.
```

### Strings: Special Case
#important 

- Strings are **value types but behave like references**.
	- Underlying data is immutable
	- Internally: pointer + length
	- Safe to share

```go
s1 := "hello"
s2 := s1

s2 = "world"
fmt.Println(s1) // hello
// No mutation possible.
```

### Function Arguments

- The REAL RULE - `Everything in Go is passed by value.`
- The difference is **what is being copied**.

#### Passing Value Type

```go
func f(x int) {
    x = 100
}

a := 10
f(a)
fmt.Println(a) // 10
// value is copied
```

#### Passing Reference Type

```go
func f(s []int) {
    s[0] = 100
}

a := []int{1, 2, 3}
f(a)
fmt.Println(a) // [100 2 3]
// The slice header was copied, but it points to the same data.
```

#### Changing Length Does NOT Affect Caller

```go
func f(s []int) {
    s = append(s, 10)
}

a := []int{1, 2}
f(a)
fmt.Println(a) // [1 2]
// Header change is local. append created a new array.
```

### Summary
#important 

| Type           | Value / Reference | Copy Behavior      |
| -------------- | ----------------- | ------------------ |
| `int`, `float` | Value             | Full copy          |
| `array`        | Value             | Full copy          |
| `struct`       | Value             | Full copy          |
| `string`       | Value (immutable) | Header copy        |
| `slice`        | Reference         | Header copy        |
| `map`          | Reference         | Pointer copy       |
| `channel`      | Reference         | Pointer copy       |
| `pointer`      | Reference         | Address copy       |
| `interface`    | Reference-like    | (type, value) copy |

