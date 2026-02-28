## Conditionals

### if/else statements

- Curly braces `{}` are **mandatory** (unlike JS).    
- No parentheses around condition.

```go
if x > 10 {
    fmt.Println("x is greater than 10")
} else if x == 10 {
    fmt.Println("x is 10")
} else {
    fmt.Println("x is less than 10")
}

// short hand
if n := len(name); n > 5 {
    fmt.Println("long name")
} else {
    fmt.Println("short name")
}
// n is scoped to the if/else block only
```

### Switch Case

- Only one case executes at a time. It does not trickle though all cases until `break` (unlike C/JS).

```go
// Basic case
switch day {
case "Mon":
    fmt.Println("Start of week")
case "Fri":
    fmt.Println("Weekend is near")
default:
    fmt.Println("Midweek")
}

// multiple cases in one go
switch num {
case 1, 3, 5, 7:
    fmt.Println("Odd")
case 2, 4, 6, 8:
    fmt.Println("Even")
}

// Switch Without Condition
score := 75
switch {
case score >= 90:
    fmt.Println("A")
case score >= 75:
    fmt.Println("B")
default:
    fmt.Println("C")
}

// Acts like a cleaner version of multiple `if/else if`.
```
#### Fallthrough 

- Once a case matches, execution stops unless you explicitly use `fallthrough`.
- `fallthrough` **forces execution to continue into the next case**, **ignoring its condition**.

```go
switch x := 2; x {
case 2:
    fmt.Println("Two")
    fallthrough
case 3:
    fmt.Println("Three")
}

// Two
// Three

```

## Loops

- Go has **only `for`** loop, used in multiple ways.
### Basic Loops 

```go
// Classic For Loop
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

// While-Style Loop
i := 0
for i < 5 {
    fmt.Println(i)
    i++
}

// Infinite Loop
for {
    fmt.Println("forever")
    break
}
```

### `for ... range` (Collections)

```go
// Slice/Array
nums := []int{10, 20, 30}
for i, v := range nums {
    fmt.Println(i, v)
}

// Map
m := map[string]int{"a": 1, "b": 2}
for k, v := range m {
    fmt.Println(k, v)
}

// String (Unicode-safe)
for i, r := range "GoLang" {
    fmt.Println(i, string(r))
}
```

## Loop Control Statements

- **break** → exits loop. 
- **continue** → skips to next iteration.
- **goto** → jumps to labeled statement (rarely used).

```go
for i := 0; i < 5; i++ {
    if i == 2 {
        continue
    }
    if i == 4 {
        break
    }
    fmt.Println(i)
}

// Labels with Loops
// break outer breaks outer loop, not just inner.
outer:
for i := 1; i <= 3; i++ {
    for j := 1; j <= 3; j++ {
        if i*j > 3 {
            break outer
        }
        fmt.Println(i, j)
    }
}
```

## Some Use Cases

### Unicode-Safe String Iteration

Strings in Go are stored as **UTF-8 encoded bytes**.
- Using a **normal for loop** (`for i := 0; i < len(s); i++`) iterates byte-by-byte.
- Some characters (like `é`, `😊`) take **multiple bytes** in UTF-8.    
- This can cause issues if you try to process text by bytes instead of runes (Unicode code points).

1. **Example (without `range`)**
```go
s := "Gó😊"
for i := 0; i < len(s); i++ {
    fmt.Printf("%d: %c\n", i, s[i]) // prints raw bytes
}

// output: 
// 0: G
// 1: Ã
// 2: B
// 3: ð
// 4: // garbled because multi-byte chars split incorrectly
// 5: �
```

2. **Example (with `range`):**
```go
for i, r := range "Gó😊" {
    fmt.Printf("%d: %c\n", i, r)
}

/*
// output

0: G
1: ó
3: 😊
*/
```

- `range` converts each sequence of bytes into a proper **rune** (int32 → Unicode code point).  
That’s what “Unicode-safe” means.

### Switch Control Inside Loops

1. `break` inside switch breaks **out of the switch only**, not the loop.
	```go
	for i := 0; i < 5; i++ {
	    switch i {
	    case 2:
	        fmt.Println("Found 2, breaking switch")
	        break // exits switch, not loop
	    default:
	        fmt.Println("i =", i)
	    }
	}
	
	/*
	output 
	
	i = 0
	i = 1
	Found 2, breaking switch
	i = 3
	i = 4
	*/
	
	// If you want escape out of the loop - use label
	outer: // label
	for i := 0; i < 5; i++ {
	    switch i {
	    case 2:
	        fmt.Println("Breaking out of the loop at i=2")
	        break outer // exits the loop, not just switch
	        //continue outer // works with continue too
	    default:
	        fmt.Println("i =", i)
	    }
	}
```

2. `continue` inside switch skips to the **next iteration of the loop**.
	```go
	for i := 0; i < 5; i++ {
	    switch i {
	    case 2:
	        fmt.Println("Skipping rest for i=2")
	        continue
	    default:
	        fmt.Println("i =", i)
	    }
	}
	
	/*
	output
	
	i = 0
	i = 1
	Skipping rest for i=2
	i = 3
	i = 4
	*/
	```

2. `return` inside switch exits the entire function immediately.
```go
func demo() {
    for i := 0; i < 5; i++ {
        switch i {
        case 3:
            fmt.Println("Returning at i=3")
            return
        }
        fmt.Println("i =", i)
    }
}
demo()

/*
output

i = 0
i = 1
i = 2
Returning at i=3
*/
```