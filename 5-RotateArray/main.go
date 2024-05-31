package main

import (
	"fmt"
	"github.com/soyhouston256/algorithms/5-RotateArray/algorithm"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("After", nums)
	algorithm.Rotate(nums, 3)
	fmt.Println("Before", nums)

	nums2 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("After", nums2)
	algorithm.Rotate2(nums2, 3)
	fmt.Println("Before", nums2)
}
