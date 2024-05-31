package main

import (
	"fmt"
	"github.com/soyhouston256/algorithms/3-MoveCeros/algorithm"
)

func main() {
	nums := []int{-2, 1, -3, 0, -1, 2, 1, -5, 4}
	fmt.Println("After", nums)
	algorithm.RemoveZeros(nums)
	fmt.Println("Before", nums)
}
