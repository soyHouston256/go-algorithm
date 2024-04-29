package main

import "fmt"

func main() {

	arr := "652notsuohyos si eman ym iH"
	fmt.Println(reverse(arr))
	fmt.Println(reverse2(arr))
}

func reverse(str string) string {
	var reversed string
	for i := len(str) - 1; i >= 0; i-- {
		reversed += string(str[i])
	}
	return reversed
}

func reverse2(str string) string {
	var reversed string
	for _, char := range str {
		reversed = string(char) + reversed
	}
	return reversed
}
