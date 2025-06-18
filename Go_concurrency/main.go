package main

import (
	"fmt"
	"time"
)

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(40 * time.Millisecond)
		fmt.Print(i, " ")
	}
}
func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(25 * time.Millisecond)
		fmt.Print(string(i), " ")
	}
}
func Calsquares(squchannel chan int, n int) {
	sum := 0
	for n != 0 {
		digit := n % 10
		sum += digit * digit
		n /= 10
	}
	squchannel <- sum
}
func Calcubes(cubchannel chan int, n int) {
	sum := 0
	for n != 0 {
		digit := n % 10
		sum += digit * digit * digit
		n /= 10
	}
	cubchannel <- sum
}
func GoRoutineWithChannels() {
	fmt.Println("--------Go routine with channels ---------")
	squchannel := make(chan int)
	cubchannel := make(chan int)

	go Calsquares(squchannel, 5)
	go Calcubes(cubchannel, 5)
	squareVal, cubVal := <-squchannel, <-cubchannel

	fmt.Println("\nCalculating Sum of squares and cubes of digits in 5", squareVal+cubVal)

}

func main() {
	go numbers()
	go alphabets()

	GoRoutineWithChannels()

	fmt.Println("Trying to go forward")

	// 	// Let Go use multiple OS threads (equal to number of logical CPUs)
	// 	runtime.GOMAXPROCS(3)

	time.Sleep(3000 * time.Millisecond) // Wait for goroutines to finish
	fmt.Println("\nMain function finished executing")

	//Channels *****************
	// var a chan int
	a := make(chan int) // Unbuffered channel with capacity 0
	go func() {
		val := <-a // Receive value from channel
		fmt.Println("Received value:", val)
	}()
	a <- 122 // Send value to channel (now there is a receiver)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Main function done with channel operations\n")

	// *****************
	x := make(chan string, 2) // Buffered channel with capacity 2
	x <- "Hello"              // sender
	x <- "Vinoth"
	str := <-x // receiver
	str1 := <-x
	fmt.Println("Data received from buffered channel: ", str)
	fmt.Println("Data received from buffered channel: ", str1)
}
