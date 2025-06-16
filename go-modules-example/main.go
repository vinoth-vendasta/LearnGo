package main

import (
	"fmt"
	"github.com/vinoth-vendasta/go-modules-example/mathutils"
)


func main() {
	var str string = "Hello, Vinoth Kumar N"
	fmt.Println(str)

 	// Using the mathutils package
	fmt.Println("Addition: ", mathutils.Add(100, 200))
}