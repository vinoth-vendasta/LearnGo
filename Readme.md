# Basics of Golang

## Data types
```golang
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```
## methods

In Go, a method is just a function with a **receiver** — a special argument that binds the function to a type (usually a struct).

## Interface
-- An interface type is defined as a set of method signatures.

**Why Important**
* Interfaces define behavior in Go. They allow you to write flexible, modular, and testable code by decoupling implementation from usage.
* Go’s interfaces support polymorphism (many types, one behavior).
* Go uses implicit satisfaction (no need to declare "implements").
* Interfaces enable abstraction and dependency injection.

### Interface Satisfaction (Implicit)
-- You don’t declare that a type implements an interface. Go checks at compile time whether a type matches an interface.

### Type Assertion and Type Switch
✅ Type Assertion
```golang
x := interface{}(42)
i := x.(int)         // assert it's an int
```
✅ Type Switch
```golang
x := 12
switch v := x.(type) {
case int:
    fmt.Println("It's an int:", v)
case string:
    fmt.Println("It's a string:", v)
default:
    fmt.Println("Unknown type")
}
```

### Interface Composition
-- Interfaces can embed other interfaces.

```golang
type Reader interface {
    Read(p []byte) (n int, err error)
}
type Writer interface {
    Write(p []byte) (n int, err error)
}
type ReadWriter interface {
    Reader
    Writer
}
```