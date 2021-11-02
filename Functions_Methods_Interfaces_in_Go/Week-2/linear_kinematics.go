package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return (0.5 * a * math.Pow(t, 2)) + (v0 * t) + s0
	}
}

func main() {
	var (
		acc      float64
		initVel  float64
		initDisp float64
		time     float64
	)

	fmt.Println("Please enter the acceleration")
	fmt.Scan(&acc)
	fmt.Println("Please enter the initial velocity")
	fmt.Scan(&initVel)
	fmt.Println("Please enter the initial displacement")
	fmt.Scan(&initDisp)

	fmt.Println("Generating function for the given parameters...")
	fn := GenDisplaceFn(acc, initVel, initDisp)

	fmt.Println("Please enter the time")
	fmt.Scan(&time)

	fmt.Printf("The final displacement is: %v\n", fn(time))
}
