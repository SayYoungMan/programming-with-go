package main

import (
	"fmt"
)

func main() {
	var floating_number float64
	var converted int64

	fmt.Scan(&floating_number)
	converted = int64(floating_number)
	fmt.Printf("%d\n", converted)
}
