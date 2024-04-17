package main

import (
	"fmt"
	"github.com/soyhouston256/algorithms/findToSum/findCommondItem"
)

func main() {

	array := []string{"a", "b", "c", "d", "e"}
	array2 := []string{"f", "g", "h", "i", "e"}
	fmt.Println(findCommondItem.FindCommonItem(array, array2))
}
