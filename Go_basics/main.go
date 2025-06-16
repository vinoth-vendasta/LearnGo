package main

import (
	"errors"
	"fmt"
	"learngo/go_basics/math"
	"time"
)

func varTypesConst() {
	var str string = "Hello, Vinoth"
	var boolean bool = true
	// byte := byte(100)
	// fmt.Println("Byte: ", byte)
	fmt.Println("\n---Variables, Types amd Constants-------------")
	fmt.Println(str)
	fmt.Println(boolean)
}

func ForIfSwitchDefer() {
	fmt.Println("\n----For, If, Switch and Defer-------------")
	for i := range 5 {
		if i == 3 {
			fmt.Println("i is 3")
		}
	}
	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning")
	case t.Hour() < 18:
		fmt.Println("Good Afternoon")
	case t.Hour() < 22:
		fmt.Println("Good Evening")
	default:
		fmt.Println("Good Night")
	}
	defer fmt.Println("defer: This will be printed last")
	fmt.Println("Current time is:", t.Format("15:04:05"))
}

func checkError(msg string) (string, error) {
	fmt.Println("\n----Error Handling-------------")
	if msg == "" {
		return "", errors.New("message cannot be empty")
	}
	return "Hello, " + msg, nil
}

func main() {
	var str string = "Hello, Vinoth"
	fmt.Println(str)
	fmt.Println("Addition: :", math.Add(100, 200))
	varTypesConst()
	ForIfSwitchDefer()

	// Error Handling
	if msg, err := checkError(""); err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(msg)
	}
}
