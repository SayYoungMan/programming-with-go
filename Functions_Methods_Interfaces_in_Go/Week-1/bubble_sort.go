package main

import "fmt"

func Swap(nums []int, i int) {
	nums[i-1], nums[i] = nums[i], nums[i-1]
	return
}

func BubbleSort(nums []int) {
	for x := 0; x < len(nums); x++ {
		for i := 1; i < len(nums); i++ {
			if nums[i-1] > nums[i] {
				Swap(nums, i)
			}
		}
	}

	return
}

func main() {
	var (
		input int
		nums  []int
	)

	for i := 0; i < 10; i++ {
		fmt.Printf("Print enter an integer (%d / 10)\n", i+1)
		fmt.Scan(&input)
		nums = append(nums, input)
	}

	BubbleSort(nums[:])

	fmt.Printf("%v\n", nums)
}
