package main

import (
	"fmt"
	"sort"
)

func concurrentSort(arr []int, c chan []int) {
	sort.Ints(arr)
	c <- arr
}

func mergeInOrder(left, right []int) []int {
	rv := make([]int, len(left)+len(right))
	i, lp, rp := 0, 0, 0

	for lp < len(left) && rp < len(right) {
		if left[lp] > right[rp] {
			rv[i] = right[rp]
			rp++
		} else {
			rv[i] = left[lp]
			lp++
		}
		i++
	}

	for lp == len(left) && rp < len(right) {
		rv[i] = right[rp]
		rp++
		i++
	}

	for lp < len(left) && rp == len(right) {
		rv[i] = left[lp]
		lp++
		i++
	}

	return rv
}

func main() {
	var (
		array      []int
		num        int
		sectionLen int
	)

	for {
		fmt.Println("Please enter a integer number. Input -9999 when you are finished entering.")
		fmt.Scan(&num)
		if num == -9999 {
			break
		}
		array = append(array, num)
	}

	sectionLen = len(array) / 4

	c := make(chan []int, 4)
	go concurrentSort(array[:sectionLen], c)
	go concurrentSort(array[sectionLen:sectionLen*2], c)
	go concurrentSort(array[sectionLen*2:sectionLen*3], c)
	go concurrentSort(array[sectionLen*3:], c)

	slice1 := <-c
	slice2 := <-c
	slice3 := <-c
	slice4 := <-c

	fmt.Printf("The result array is: %v", mergeInOrder(mergeInOrder(slice1, slice2), mergeInOrder(slice3, slice4)))
}
