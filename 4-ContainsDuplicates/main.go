package main

import (
	"fmt"
	"github.com/soyhouston256/algorithms/4-ContainsDuplicates/algorithm"
)

func main() {
	nums := []int{-2, 9, -3, 7, 1, 2, 1, -5, 4}
	fmt.Println("result", algorithm.ContainsDuplicate(nums))
}
