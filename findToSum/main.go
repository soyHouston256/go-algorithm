package main

import (
	"fmt"
	"github.com/soyhouston256/algorithms/findToSum/myArray"
)

func main() {

	arr := myArray.NewArray()

	arr.Push("a")
	arr.Push("e")
	arr.Push("i")
	arr.Push("o")
	arr.Push("u")

	val, _ := arr.Get(0)
	fmt.Println(val)
	fmt.Println(arr)
	arr.Delete(4)
	fmt.Println(arr)

}
