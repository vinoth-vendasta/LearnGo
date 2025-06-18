# 🧵 Go Compilation Process – Key Notes

This document explains the step-by-step compilation process of Go source code into a native executable binary.

---

## 📄 Step 1: Write Go Code

```go
// main.go
package main

import "fmt"

func main() {
	fmt.Println("Hello, Go!")
}
```

## 📄 Step 2: Use Go Tools

| Command      | Description                          |
| ------------ | ------------------------------------ |
| `go run`     | Compiles & runs the code temporarily |
| `go build`   | Compiles code into a binary          |
| `go install` | Compiles and installs to `$GOBIN`    |
| `go test`    | Compiles & runs tests                |

go build main.go

## 🔍 Step 3: Lexical Analysis & Parsing

* Tokenizes source code (identifiers, keywords, symbols)
* Parses into an AST (Abstract Syntax Tree)
* Performs syntax and type checking

## 🧠 Step 4: Intermediate Representation (IR)
* Translates AST → SSA IR (Static Single Assignment)
* Performs optimizations:
    - Dead code elimination
    - Escape analysis
    - Inlining
    - Bounds checking
    - Constant folding

## 🏗️ Step 5: Code Generation
* Converts optimized IR into platform-specific machine code
* Handles CPU architecture & OS targeting

## 🔗 Step 6: Linking
* Combines all .o object files into one binary
* Resolves cross-package references
* Links Go runtime and optional C libraries (cgo)

## 📦 Step 7: Output Binary
```go
./main
# Output: Hello, Go!
```
* A fully self-contained native executable is produced. No external runtime needed!
