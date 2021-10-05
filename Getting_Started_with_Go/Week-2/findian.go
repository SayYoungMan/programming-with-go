package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	fmt.Scan(&input)

	if strings.ContainsAny(input, "ian") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
