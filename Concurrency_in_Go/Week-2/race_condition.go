package main

import (
	"fmt"
	"sync"
)

var (
	x  int = 0
	wg sync.WaitGroup
)

func read() {
	for x < 100 {
		var readValue = x
		fmt.Printf("Current value read: %v \n", readValue)
	}
	fmt.Println("Reading done!")
	wg.Done()
}

func update() {
	for x < 100 {
		x += 1
		fmt.Printf("Current value written: %v \n", x)
	}
	fmt.Println("Updating done!")
	wg.Done()
}

func main() {

	wg.Add(2)

	go read()
	go update()

	wg.Wait()
}
