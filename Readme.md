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
***
## Methods :
In Go, a method is just a function with a **receiver** — a special argument that binds the function to a type (usually a struct).
***

## Interface :
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
- Interfaces can embed other interfaces.

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
***
## Concurreny in Go (eg. capable of handling both)
 * Let’s consider a person jogging. During his morning jog, let’s say his shoelaces become untied. Now the person stops running, ties his shoelaces and then starts running again. This is a classic example of concurrency

 **parallelism** 
 * Let’s understand it better with the same jogging example. In this case, let’s assume that the person is jogging and also listening to music on his iPod. In this case, the person is jogging and listening to music at the same time, that is he is doing lots of things at the same time. This is called parallelism.

 ## Go Routines
 * Goroutines can be thought of as lightweight threads. The cost of creating a Goroutine is tiny when compared to a thread

    **How to start a Goroutine?**
    - Prefix the function or method call with the keyword go and you will have a new Goroutine running concurrently.

    **properties of goroutines**
    - When a new Goroutine is started, the goroutine call returns immediately. Unlike functions, the control does not wait for the Goroutine to finish executing. The control returns immediately to the next line of code after the Goroutine call and any return values from the Goroutine are ignored.
    - The main Goroutine should be running for any other Goroutines to run. If the main Goroutine terminates then the program will be terminated and no other Goroutine will run.
    
    * function main() is the "main goroutine" which runs all the goroutines.
    * If main goroutine terminated, the program will stop the execution

### Go Channels
* Channels can be thought of as pipes using which Goroutines communicate
* 

```go
package main

import "fmt"

func main() {
	var a chan int
	if a == nil {
		fmt.Println("channel a is nil, going to define it")
		a = make(chan int)
		fmt.Printf("Type of a is %T", a)
	}
}
```
**Sends and receives are blocking by default if we use Channels**
* No need for explicit blocking or conditional Statements

```go
package main

import (
	"fmt"
)
func hello(done chan bool) {
	fmt.Println("Hello world goroutine")
	done <- true // sending
}
func main() {
	done := make(chan bool)
	go hello(done)
	<-done // receiving
	fmt.Println("main function")
}
```
### Deadlocks
* One important factor to consider while using channels is deadlock. If a Goroutine is sending data on a channel, then it is expected that some other Goroutine should be receiving the data. If this does not happen, then the program will panic at runtime with Deadlock.
```go
package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 2) //capacity 2 {store 2 values without blocking}
	ch <- "naveen"
	ch <- "paul"
	ch <- "steve" //sending to channel 
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// no receiver for overflow sender
}
```

```console
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
	/tmp/sandbox046150166/prog.go:6 +0x50
```

### Undirectional Channels
* It only send or receive data
```go
package main

import "fmt"

func sendData(sendch chan<- int) {
	sendch <- 10
}

func main() {
	sendch := make(chan<- int)
	go sendData(sendch)
	fmt.Println(<-sendch)
}
```
### WaitGroup
*  WaitGroup is used to wait for a collection of Goroutines to finish executing
- When we call Add on the WaitGroup and pass it an int, the WaitGroup’s counter is incremented by the value
- The Wait() method blocks the Goroutine in which it’s called until the counter becomes zero.
- The way to decrement the counter is by calling Done() method on the WaitGroup

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done()
}

func main() {
	no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1)
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")
}
```
