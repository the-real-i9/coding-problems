package main

import (
	"coding/problemSet1"
	"fmt"
)

func main() {

	medianOfTwoSortedArrays := problemSet1.MedianOfTwoSortedArrays

	fmt.Print("medianOfTwoSortedArrays([]int{1, 3}, 2, []int{2}, 1): ")
	fmt.Println(medianOfTwoSortedArrays([]int{1, 3}, 2, []int{2}, 1)) // Answer: 2

	fmt.Print("medianOfTwoSortedArrays([]int{1, 3}, 2, []int{2, 4}, 2): ")
	fmt.Println(medianOfTwoSortedArrays([]int{1, 3}, 2, []int{2, 4}, 2)) // Answer: 2.5

	fmt.Print("medianOfTwoSortedArrays([]int{1}, 1, []int{2, 4}, 2): ")
	fmt.Println(medianOfTwoSortedArrays([]int{1}, 1, []int{2, 4}, 2)) // Answer: 2

	fmt.Print("medianOfTwoSortedArrays([]int{2,4}, 2, []int{2, 4}, 2): ")
	fmt.Println(medianOfTwoSortedArrays([]int{2, 4}, 2, []int{2, 4}, 2)) // Answer: 3

	fmt.Print("medianOfTwoSortedArrays([]int{2,4}, 2, []int{}, 0): ")
	fmt.Println(medianOfTwoSortedArrays([]int{2, 4}, 2, []int{}, 0)) // Answer: 3
}
