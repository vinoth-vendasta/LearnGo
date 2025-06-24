package main

import (
	"fmt"
	"sync"
)

var x = 0

func Increment(wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock()
	x = x + 1
	mu.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	for range 100 {
		wg.Add(1)
		go Increment(&wg, &mu)
		fmt.Println("Current value of x: ", x)
	}
	wg.Wait()
	fmt.Println("Final value of x: ", x)
}
