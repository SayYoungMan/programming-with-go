package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type person struct {
		Name    string
		Address string
	}

	var (
		n string
		a string
	)

	fmt.Printf("Please enter a name of a person\n")
	fmt.Scan(&n)

	fmt.Printf("Please enter an address of %s\n", n)
	fmt.Scan(&a)

	p := person{
		Name:    n,
		Address: a,
	}

	b, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		fmt.Println("JSON Marshalling failed")
	}

	fmt.Println(string(b))
}
