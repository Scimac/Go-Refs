## Go Commands

- `go run` - complies and executes the bin. No output.
- `go build` - compiles the project and returns the executable file.

## Go variables

- Declaration is done as - 

```go
// `var` declares 1 or more variables.
var a = "initial"

// You can declare multiple variables at once.
var b, c int = 1, 2

// Go will infer the type of initialized variables.
var d = true

// Variables declared without a corresponding initialization are _zero-valued_. 
// For example, the zero value for an `int` is `0`.
var e int

// The `:=` syntax is shorthand for declaring and initializing a variable,
// This syntax is only available inside functions.
f := "apple"   // means var f string = "apple"   
```

## Functions

Functions consists of - 

```go
// keyword -- func name -- arguments -- return type
func newCard() string {
	return "Five of Spades"
}
```

## Arrays and Slices

- Arrays are fixed length data structure of same datatype.
- Slice are variable length data structure of same datatype.

```go
// Array declaration
var a [5]int
a[4] = 100

b := [5]int{1, 2, 3, 4, 5}

// slice declaration
var s []string

// make function is an inbuild function
s = make([]string, 3)

// can be declared just like an array.
b := []int{1, 2, 3, 4, 5}

// append function is used to append the data into the card
// causes deep copy or creates an new array/slice
cards = append(cards, 'Six of Spades')
```

## For loop

```go
// index and card are variables for iterating
// := operator used as they get destroyed in each iteration
for index,card := range cards {
	fmt.PrintLn(index,card)
}
```

## Go as OO language

- Go is not a OO language. But can be made to work like one.
- So, we have 5 base types in Go - 
	- integer
	- float
	- string
	- array
	- map
- To create a new object or type in Go, we extend a base type and add functionality to it.

### Example

- We can create a type as `deck` and add functions with `deck` as the receiver to it.

- Receiver functions in Go are functions that are associated with a type.
- In this case, the `print` function is associated with the `deck` type.
- `d` in the function is the actual instance of the `deck` type.
- By convention, receiver names are usually short and descriptive.
- **They are often single letters or short words that indicate the type of the receiver.**

>  Treat d as this in JS.

```go
// deck.go

import 'fmt'

// extension of a type
type deck []string

// Functions with deck as a receiver

// print prints the contents of the deck
// print is a method oftype deck
func (d deck) print() {
	for i, card := range d {
		// Print the index and the card
		// Using fmt.Printf for formatted output
		// %v is used to print the value of the card
		// %d is used to print the index
		fmt.Printf("%d: %v\n", i, card)
	}
}

// deal takes a deck and a hand size, and returns two decks:
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
```

## Slices in GO

- Slices in Go are 0-indexed and behave like an growing array.
- Indexing works exactly like any other programming language.
- Slicing works like python in Go.

```go
type fruits []string

// explain indexing and sub-slicing
func print(f fruits) {
	fals := fruits{}
	fals = append(fals, "apple", "banana", "orange", "grape", "kiwi", "mango", "peach", "pear", "pineapple", "strawberry")

	// indexing examples in Go          
	
	// Accessing the first element
	fmt.Println("First fruit:", fals[0])       
	// Accessing the last element   
	fmt.Println("Last fruit:", fals[len(f)-1])       
	// Slicing to get the first three elements 
	fmt.Println("First three fruits:", fals[:3])   
	// Slicing to get the last three elements
	fmt.Println("Last three fruits:", fals[len(f)-3:]) 
	// Accessing the middle element
	fmt.Println("Middle fruit:", fals[len(fals)/2])                 
	// Printing the entire slice  
	fmt.Println("All fruits:", fals)    

	// examples of slicing
	// Slicing to get the first five elements
	fmt.Println("Slicing the first five fruits:", fals[:5]) 
	// Slicing to get the last five elements
	fmt.Println("Slicing the last five fruits:", fals[len(fals)-5:]) 
	// Slicing from index 2 to 5
	fmt.Println("Slicing from index 2 to 5:", fals[2:5]) 
	// Slicing from index 3 to the end
	fmt.Println("Slicing from index 3 to the end:", fals[3:])  
	// Slicing from the start to index 4
	fmt.Println("Slicing from the start to index 4:", fals[:4])
	// Slicing the entire slice	
	fmt.Println("Slicing the entire slice:", fals[:]) 
}

```

## Byte Slices

- A byte slice is a array representation of bytes corresponding to the actual data.
- Represented as - `[]byte`

```Go
// converting the deck type into strings then to byte strings
// deck --> []string --> string --> []byte

import (
	"fmt"
	"strings" // standard package can be found in Go docs
)

type byte []string

func (d deck) toString() string {
	// deck to string slice
	// []string(d) -- snippet for type casting
	
	// Let's convert the slice of strings to a string type
	return strings.Join([]string(d), ',')
}

// for converting to the byte string
// []byte(cardsString)]
```

### Reading and writing the files

- Golang uses the `ioutils` repo for completing the file read write operations..

> `ioutils` is deprecated, now all methods are moved to 'os' library.

```go
// earlier it was - import "io/ioutils"
import "os"

// For saving the data to the file.
func (d deck) saveToFile(filename string) error {
	cardsString := d.toString()

	// The third argument (0666) sets the file permissions.
	// Common permissions:
	// 0644 - Owner can read/write, others can read
	// 0600 - Owner can read/write, others have no permissions
	// 0666 - Everyone can read/write (not recommended for sensitive files)
	// 0755 - Owner can read/write/execute, others can read/execute (for executables)
	err := os.WriteFile(filename, []byte(cardsString), 0644)
	// earlier it was err := ioutils.WriteFile(...)
	
	return err
}
```

### Random Number generation

- Idea for shuffling is to swap the current card with a card at random index. This is enough to create a confusion of initial indexes.
- `Intn` function from `math/rand` is used to generate the random integer.

```go
// shuffle function
func (d deck) shuffle() {
	// Implementing a simple shuffle algorithm
	for i := len(d) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		d[i], d[j] = d[j], d[i] // Swap cards
	}
}
```

- Random library uses a **seed value** for creating the random values.
- If the seed value remains constant between the execution cycles, we may see similar patterns instead of absolute randomness.
#### Creating new Rand type with custom seed

- Let's see how to create a new `rand`, instead of using the default one and pass the seed value to it. 
- Then use the `intn` method of that new `rand` to generate random number.
- Use doc - https://pkg.go.dev/math/rand#Rand

```go
// shuffle with custom random seed
func (d deck) shuffleWithSeed() {
	// Implementing a simple shuffle algorithm with a custom seed
	// using current time in nanoseconds as a seed 
	n := time.Now().UnixNano() // used this as seed needs to be int64

	seededRand := rand.New(rand.NewSource(n)) // Example seed

	for i := len(d) - 1; i > 0; i-- {
		j := seededRand.Intn(i + 1)
		d[i], d[j] = d[j], d[i] // Swap cards
	}
}
```

### Testing in Go

> we need `go.mod` file while running tests. can be created using `go mod init cards`
> see 'go help modules'

- All the testing functions are called using a first parameter of type `*testing.T` - this is called a testing handler object.
- The names of the testing functions usually starts with capital first letter, viz. testing function for `newDeck` is named as `TestNewDeck`.
- Follow the convention, so that we know which functions are mapped to which tests.

```go
func TestSaveToFileAndLoadFromFile(t *testing.T) {
	// clean up before and after the test
	// always name the temp file something unique to avoid conflicts
	os.Remove("_decktesting")

	d := newDeck()
	err := d.saveToFile("_decktesting")

	if err != nil {
		t.Errorf("Error saving deck to file: %v", err)
	}

	loadedDeck, err := loadFromFile("_decktesting")

	if err != nil {
		t.Errorf("Error loading deck from file: %v", err)
	}

	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length of 52, but got %d", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
```

### Structs

- Structs in GoLang are collection named elements, called fields, each of which has a name and a type.
- They can be used to group similar data together. They are analogous to `objects` in JS, or `Dictionaries` in Python.
#### Defining a struct 

```go
// Create a new struct type
	type Person struct {
		FirstName string
		LastName  string
		Age       int
	}
```

#### Initialising a struct, basic operations

- There are roughly **three** ways to initialise a struct - 
```go
	// First is - Can be used when values are not previously known
	// Create a new instance of the struct
	var p Person
	
	// Assign values to the fields
	p.FirstName = "John"
	p.LastName = "Doe"
	p.Age = 30
	
	// Second is - The most reliable way of declaring the struct
	p2 := Person{
		FirstName: "Jane",
		LastName: "Smith",
		Age: 25,
	}
	
	// Third is - Short hand and most error prone way
	p3 := Person{"Alice", "Johnson", 28}
```

#### Embedding a struct

- Embedding structs means using a struct as a field in another struct.
```go
	type Adress struct {
		PinCode int
		City    string
		Country string
		Email   string
	}

	// There is a short hand for embedding a struct. 
	// The name and type of the can be written as a single thing.
	type Employee struct {
		Person   // Embedded struct
		Position string
		Adress   // Embedded struct
	}

	func main() {
		e := Employee{
			Person: Person{
				FirstName: "Bob",
				LastName:  "Brown",
				Age:       35,
			},
			Position: "Manager",
			Adress: Adress{
				PinCode: 12345,
				City:    "New York",
				Country: "USA",
				Email:   "bob.brown@mail.com", // commas are needed on all lines
			},
		}
	}
```

#### Receiver functions for structs

- They are written similar to the other DS.
```go
// employee introducing methods
func (e Employee) Introduce() string {
	return "Hello, my name is " + 
		e.FirstName + " " + 
		e.LastName + ", I am a " + 
		e.Position + " living in " + 
		e.City + ". You can reach me at " + e.Email
} 

// person birthday method
func (p *Person) Birthday() {
	p.Age++
}
```

### Pointers in Go

- GoLang is a **pass by value** language.
- So, whenever we write a receiver function for a type, (types of Data Structures that are `value` type, not `reference` type), Go creates a copy of the value and passes that value as a argument to the receiver function.
- So, for `value` type data structures, the original value of the DS remains untouched in the receiver function scope.

```go
package main

import (
	"fmt"
)

type Person struct {
	name string
	lastName string
	age int
	email string
}

func main () {
	jim := Person{
		name: "Jim",
		lastName: "Anderson",
		age: 25,
		email: "jim.andy@gmail.com", // don't forge the trailng comma
	}
	
	jim.updateName("Jimmy")
	jim.print()
}

func (p Person) print() {
	fmt.Printf("%+v", p)
}

func (p Person) updateName(newName string) {
	p.name = newName
}

// output is - 
// {name:Jim lastName:Anderson age:25 email:jim.andy@gmail.com}

// as opposed to -
// {name:Jimmy lastName:Anderson age:25 email:jim.andy@gmail.com}% 
```

- Pointers come in handy here. 
- Instead of passing the value that gets copied, we pass the `address of the value` of the DS instead. ( which doesn't gets copied )
- Pointers are nothing but address holders that "point" to the memory address that holds the actual value of the DS.

```go
// The little change that is needed is - 

{ // ...
	
	jimPointer := &jim // Pointer created with address to value
	jimPointer.updateName("Jimmy")
	
	// this remains jim as not update to value is done here
	jim.print()
}

// here *Person refers to the type of the pointer
func (pointerToPerson *Person) updateName() {
	// here we access the value at the address held by pointer
	// using *pointerToPerson
	(*pointerToPerson).name = updatedName
}

// now the output is - 
// {name:Jimmy lastName:Anderson age:25 email:jim.andy@gmail.com}
```

- `&variableName` - means fetching address of the memory location where the value of DS is placed.
- `*pointerVariable` - means accessing the value from the address held by the pointer, i.e the memory location where the original value of DS is placed. ( This is called the dereferencing the pointer )

- `*variable` can have two meanings in Go - 
	- `*pointer` - i.e * in front of pointer variable that holds address - means dereferencing the pointer. E.g. *pointerToPerson
	- `*typeVariable` - i.e * in front of variable where `type` of variable is defined - means we are working with a pointer variable of that particular type. E.g. `*Person`

#### Shortcut to pointers

- There is a shorthand in go, for simplicity sake, to not pass the address all the time.
```go
{ // ...
	
	// jimPointer := &jim // Pointer created with address to value
	// jimPointer.updateName("Jimmy)
	
	// The above portion can be replaced by - 
	jim.updateName("Jimmy")
	jim.print()
}

func (pointerToPerson *Person) updateName() {
	(*pointerToPerson).name = updatedName
}
```

- So basically, the `value` type DS can be passed to a receiver function that expects pointer of that type. 
- `Go` internally fetches the address of the value of that DS, and passes the new pointer to the receiver function instead.
- This is a big thing to remember, that is taken care internally.

### Reference vs Value type DS in GO

- Let's do a exercise first, what is the output of this code snippet -
```go
package main

import "fmt"

type slice []string

func main() {
	mySlice := slice{"Hi", "There", "you", "GO"}
	
	mySlice.update("Bye", 0)
	fmt.Printf("%v", mySlice)
}

func (s slice) update(newWord string, pos int) {
	s[pos] = newWord
}

// Based on previous learnings, you'd expect the output to be - 
// [Hi There you GO] - since pointers not used,

// But actual output is - [Bye There you GO]
```

- The trick lies in the type of the DS - `slice` is a `referenced` DS in GO.
#### Classifications of DS in GO

|                 | Value Type                                   | Reference Type                                                |
| --------------- | -------------------------------------------- | ------------------------------------------------------------- |
| **Examples**    | `int`, `float64`, `bool`, `struct`, `array`  | `slice`, `map`, `channel`, `function`, `pointer`, `interface` |
| **Description** | Stored directly; assignment copies the value | Stores reference to data; assignment                          |

| Value Type               | Reference Type      | Description                                                 |
| ------------------------ | ------------------- | ----------------------------------------------------------- |
| `array`                  | `slice`             | Arrays are fixed-size; slices are dynamic views over arrays |
| `struct`                 | `pointer to struct` | Structs are values; pointers reference structs              |
| `int`, `float64`, `bool` | —                   | Basic value types; not typically referenced directly        |
| —                        | `map`               | Maps are reference types for key-value storage              |
| —                        | `channel`           | Channels are reference types for goroutine communication    |
| —                        | `function`          | Functions are reference types in Go                         |
#### What does a reference DS mean?

| Reference Type | Internal Storage Details                                                                                                                                                                               |
| -------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| Slice          | A slice is a small struct containing a pointer to an underlying array, a length, and a capacity. The slice itself is a value, but it references the backing array.                                     |
| Map            | A map is a pointer to a runtime hash table structure. The map variable holds a reference to this hash table, which manages buckets for key-value pairs.                                                |
| Channel        | A channel is a reference to a runtime data structure that manages a queue (buffer) and synchronization primitives for goroutine communication. The channel variable holds a pointer to this structure. |
| Function       | A function variable holds a pointer to the compiled function code and, if it's a closure, a reference to the captured environment (variables).                                                         |
| Pointer        | A pointer holds the memory address of another variable or value. It directly stores the address                                                                                                        |
- So, when passed to the receiver function, a new copy of value of referenced DS is created. But, it again in turn points to same underlying `Primitive Value Type` DS.
- Hence, we can ignore the concept of pointers when it comes to `Referenced DS`, as they themselves are pointers to their `Value Type` value holding counterparts.

- The zero value of reference types (`nil`) means they don’t point to any underlying data. Using a nil map, slice, or channel without initialisation can cause runtime errors.

- Assigning a reference type variable (like a slice or map) to another variable copies only the reference, not the underlying data. Mutations via one reference are visible to all references.


### Structs in GO

- The maps are similar to HashMaps JS or Dictionaries in Python.
- The keys of the maps should be all of one type, similarly the values of maps must also be of same type (not necessarily of the same type as the keys).

#### Declaring and initialising Maps

- The maps are declared by declaring the type of keys and values.
```go
func main () {
	// declaring and initializing a map
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	};

	// Just declaring a map, will be initalised with nil values.
	// In this case, empty map
	colors := make(map[string]string)

	// OR use builtin make function to initialize a map
	colors := make(map[string]string)

	// assign it later, if needed
	colors["white"] = "#ffffff"
	colors["black"] = "#000000"
	colors["red"] = "#ff0000"

	// iterating through the map
	for color, hex := range colors {
		println(color, hex)
	}

	// delete operations on maps
	delete(colors, "red")
	
	// println does not print the map keys in same order
	// as they were added. Maps are unordered collections
	// So the order of keys can be different each time
	// you run the program
	fmt.Println(colors)
}
```

- Dot notation cannot be used in maps, as the keys can be of different types, GO judges the type of key inside the `[]` braces. So, something like `colors.10` doesn't occur.

#### Differentiating between Maps and Structs

| Feature                | Map                                              | Struct                                         |
|------------------------|--------------------------------------------------|------------------------------------------------|
| Definition             | Collection of key-value pairs                    | Collection of named fields                     |
| Key Type               | Keys must be of a single, comparable type        | Not applicable                                 |
| Value Type             | All values must be of the same type              | Each field can have a different type           |
| Declaration            | `make(map[KeyType]ValueType)` or map literal     | `type MyStruct struct { ... }`                 |
| Access                 | Accessed by key: `m["key"]`                      | Accessed by field: `s.FieldName`               |
| Order                  | Unordered                                        | Fields have a fixed order                      |
| Mutability             | Keys and values can be added, updated, or deleted| Fields are fixed; struct values can be updated |
| Use Case               | Dynamic, unknown, or varying keys                | Fixed, known set of fields                     |
| Memory Layout          | Heap-allocated, managed by runtime               | Contiguous memory layout                       |
| Zero Value             | `nil`                                            | All fields set to their zero
### Interfaces

- Interfaces are named collections of method signatures.
- Interfaces in GO are helpful in reducing duplicating functions in the same code.
- They create a common "type" which can be "received" by the functions who share a common logic across different DS with different type, implementing same receiver attributes/functions.

### Example - 

```go
package main

// The types that render DataStructures are called - Concrete Types
// Two different bots
type englishBot struct {
	greeting string
}

type spanishBot struct {
	greeting string
}

// Interface here provides a common type, Called - Interface Type  
// Noteworthy thing is - 
// it list down all the receiver functions that are repeated between the different types
type bot interface {
	// receiver functions used by multiple structs
	// this is how we define interfaces in GO
	getGreeting() string
}

func main() {
	eb := englishBot{greeting: "Hello!"}
	sb := spanishBot{greeting: "Hola!"}

	// unnecessary duplication of logic and code
	printEnglishGreeting(eb)
	printSpanishGreeting(sb)

	// using interfaces to remove duplication of code and logic
	// since all functions in GO need to have a defined type
	// we can use interfaces to define a common type
	printGreeting(eb)
	printGreeting(sb)
}

func (eb englishBot) getGreeting() string {
	return eb.greeting
}

func (sb spanishBot) getGreeting() string {
	return sb.greeting
}

func printEnglishGreeting(eb englishBot) {
	println(eb.getGreeting())
}

func printSpanishGreeting(sb spanishBot) {
	println(sb.getGreeting())
}

// To print this greeting, we will need something like - 
//func printGreeting(eb englishBot) {
//	println(eb.getGreeting())
//}
//func printGreeting(sb spanishBot) {
//	println(sb.getGreeting())
//}

// Repeated logic can be unified
func printGreeting(b bot) {
	println("This is a interfaced function")
	println(b.getGreeting())
}
```

#### Defining Interfaces in GO

```go
// definition
type interfaceName interface {
	receiverFunctionName(receiverArgementsType) (receiverOutputType)
}

// example
type bot interface {
	getGreeting() string
	getGrammarCheck(string, int) (bool, int)
}
```

- Interfaces are implicit. Means, there is no code that ties your type to an interface. GO's compiler internally matches them.
- Interfaces are a contract between different types to share some common logic. The logic between all the receiver functions of types may not be similar.

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

### Go Routines and Channels

- Go Routines help in adding asynchronicity in the code, achieving concurrency for which Go is famous for.
#### Example

- Simple non-concurrent code snippet, making 5 series calls. 

```go
func main() {
	links := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.linkedin.com",
		"https://www.github.com",
	}

	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}
	fmt.Println(link, "is up!")
}
```

- This is resources-blocking code design, to correct this, we needs to. make all the 5 calls in parallel and handle the responses as they come back.

#### Go Routines

- **`Goroutines`** are lightweight threads managed by the Go runtime.
- You start a `goroutine` by prefixing a function call with the `go` keyword.
- They run concurrently with other `goroutines` in the same address space.

- The `main()` function itself is also a Go Routine. The Prime Go Routine.
- `go` keyword, used to spawn a Go Routine, is typically used before function calls with asynchronous elements inside them.

```go
go checkLink(link) // http.Get(link) is the blocking call 
go fmt.Println("This runs in a goroutine") // can be used like this too!!
```

- `Goroutines` are lightweight and managed by the Go scheduler.
- The scheduler multiplexes many `goroutines` onto a smaller number of OS threads.
- OS threads are then scheduled by the operating system onto available CPU cores.
- This allows Go to run thousands of `goroutines` efficiently, even on a limited number of cores.

```go
+-------------------+         +-------------------+
|   Goroutine 1     |         |   Goroutine 2     |
+-------------------+         +-------------------+
           \                         /
            \                       /
             \                     /
              \                   /
               \                 /
                \               /
                 \             /
                  \           /
                   \         /
                +-------------------+
                |   Go Scheduler    |
                +-------------------+
                   /           \
                  /             \
                 /               \
        +----------------+   +----------------+
        |   OS Thread 1  |   |   OS Thread 2  |
        +----------------+   +----------------+
                |                   |
                |                   |
        +---------------+   +---------------+
        |   CPU Core 1  |   |   CPU Core 2  |
        +---------------+   +---------------+
```

#### Concurrency vs Parallelism

1.  Concurrency

- **Definition:** Managing multiple tasks at the same time (making progress on many tasks by switching between them).
- **In Go:** Achieved using `goroutines`. The Go scheduler rapidly switches between `goroutines`, `giving` the illusion of simultaneous execution, even on a single CPU core.
- **Example:**  
   - `Goroutines` checking multiple links, but not necessarily at the exact same instant.

2.  Parallelism

- **Definition:** Actually running multiple tasks at the exact same time (simultaneously) on multiple CPU cores.
- **In Go:** If your machine has multiple cores and you set `GOMAXPROCS > 1`, `goroutines` can run in true parallel on different cores.
- **Example:**  
   - `Goroutines` checking different links at the exact same time on different CPU cores.

3. Summary Table

| Aspect   | Concurrency                         | Parallelism                     |
| -------- | ----------------------------------- | ------------------------------- |
| Meaning  | Many tasks in progress              | Many tasks running at once      |
| Hardware | 1 or more cores                     | 2 or more cores                 |
| In Go    | Goroutines (always)                 | Goroutines + multiple CPU cores |
| Analogy  | Chef cooking many dishes, switching | Many chefs cooking at once      |

- Now, making this change in our example, makes it concurrent - 
```go
	for _, link := range links {
		checkLink(link)
	}
	
	// But the output is - 
	// ... - blank
	// What happened? 
```

```go
Time →
|------------------------------------------------------>
|
| main()
|  |
|  |--+------------------- (spawns goroutines) -------------------+
|     |                    |         |         |         |        |
|     v                    v         v         v         v        v
|  go checkLink()   go checkLink()  ...   go checkLink() ... go checkLink()
|     |                    |         |         |         |        |
|     |--------------------|---------|---------|---------|--------|
|     |                    |         |         |         |        |
|     |   (main exits)     |         |         |         |        |
|     +--------------------+---------+---------+---------+--------+
|
| (child goroutines may still be running, but main exits)
```

- So, We use `Channels` for communicating between the `Goroutines`, i.e child routines telling the main routine to wait for its completion.
#### Channels

- `Channels` are **typed** pipes that allow `goroutines` to communicate and synchronise.
- You can send values into a channel from one `goroutine` and receive those values in another.
- Channels help avoid explicit locks and shared memory.

```go
ch := make(chan string)  // declaring channels with type string
ch := make(chan int) // Unbuffered channel of int
ch := make(chan string, 10) // Buffered channel of string (size 10)

ch <- value // Sending to a Channel

v := <-ch // Receiving from a Channel
v, ok := <-ch // 'ok' is false if channel is closed

close(ch) // Closing a Channel

// Range Over a Channel
for v := range ch {
    fmt.Println(v)
}

// Select Statement (for multiple channels)
select {
	case v := <-ch1:
	    // handle v from ch1
	case ch2 <- x:
	    // send x to ch2
	default:
	    // if no channel is ready
}
```

- Channels are blocking code in Go, and **routines wait for the channels to complete the data transfer before exiting the code execution**.
- If data is sent into the channel, it must be received!, and if a channel is expecting to receive a data, it must be sent. Or the routine/program hangs. 

```go
func main() {
	links := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.linkedin.com",
		"https://www.github.com",
	}

	ch := make(chan string)

	for _, link := range links {
		go checkLink(link, ch)
	}

	fmt.Println("Waiting for responses...")

	for i := 0; i < len(links); i++ {
	// for range links { // modern for loop syntax, i is not reused
		fmt.Println(<-ch)
	}
	
	// fmt.Println(<-ch) // will hang the main routine as no data is being sent
}

func checkLink(link string, ch chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		ch <- "link might be down"
		return
	}
	fmt.Println(link, "is up!")
	ch <- "link is Up"
}
```

- Lets add a polling pattern to the Site Status Checker, call must be fired again after every 5 secs.

```go
{
	// ...
	for _, link := range links {
		go checkLink(link, ch)
	}
	
	fmt.Println("Waiting for responses...")
	
	// notably, go does not have a while loop, 
	// go has an ideology of multiple syntaxes not needed for same task
	
	// This below is a true infinite while loop in go
	for {
		go checkLink(<-ch, ch) // receiving the link and respawn go routine
	}
}

func checkLink(link string, ch chan string) {
	_, err := http.Get(link)
	if err != nil {
		ch <- link // sending in the link
		return
	}
	ch <- link // sending in the link
}
```

- Alternative syntax for the `for (while)` loop- 

```go
// <-ch can be easily missed in the code
// below is much suited version of the same code

for l := range ch { // we wait here for the channels to update
		go checkLink(l, ch) // receiving the link and respawn go routine
	}
```

- For adding the 5 secs delay - 

```go
// true while loop in go
	for l := range ch {
		// function literal / anonymous function
		// l is passed as argument to avoid closure capturing the variable
		// which would lead to all goroutines using the last value of l
		// and hence all checking the same link
		// this is a common gotcha in Go
		go func(li string) {
			time.Sleep(5 * time.Second)
			checkLink(li, ch) // receive from channel and pass to goroutine
		}(l)
	}
```

- `time.Sleep()` pauses the current go routine for `x` duration of time. We could have placed it in `checkLink`, but that violates the meaning of the function. (as it is not `checklinkaftersleep` or `checklinkbeforesleep`) 
- Therefore, a function literal (anonymous functions in Go) is used as a wrapper before triggering the next call.
- So, the current go routine is "slept" before checking the status of link.

### Function literals

A **function literal** is an **anonymous function**—a function without a name. You can define it inline, assign it to a variable, or invoke it immediately.

#### Example 

```go
	// function literal
	func (li string) {
		time.Sleep(5 * time.Second)
	}
	
	// Immediately invoked
	func (li string) {
		time.Sleep(5 * time.Second)
	}(l)
```

- `func(li string) { ... }` is a function literal (anonymous function).
- `(l)` at the end **immediately invokes** the function, passing `l` as the argument for `li`.
- This pattern is often used in `goroutines` to capture the current value of a variable (here, `l`), avoiding closure pitfalls.
#### Why use function literals?

- To define short, one-off functions inline.
- To capture variables in their current state (avoiding closure bugs).
- To pass logic as arguments (like callbacks).



### Doubts

- What is a rune in go? 
	- The rune type is just int32 under the hood.
	- A string in Go is a sequence of bytes, not runes.

- What is `println`? why use `fmt.PrintLn`?
	- The difference between `println` and `fmt.Println` is that `fmt.Println` adds spaces between arguments and a newline at the end, while just prints the arguments without spaces and no newline at the end.
	- `println` is a built-in function, while `fmt.Println` is part of the `fmt` package
	- benefits of using `fmt.Println` is that it provides more formatting options and is more versatile for complex output.
	- examples of `prinln` - 
		// `println("Hello, World!")`
		// `println(42)`
		// `println(true)`
		// `println(3.14)`
		// `println("Value of p3:", p3)`
	- we even have `Printf` in `fmt` for formatted printing.
- How to call functions from another file but same package in go?
	- In Go, only identifiers (functions, types, variables) that start with an uppercase letter are exported and accessible from other packages.
	- Run both the files to get the result - `go run main.go bot.go`

