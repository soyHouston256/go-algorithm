package main

import "fmt"

func main() {

	arr := "652notsuohyos si eman ym iH"
	fmt.Println(reverse(arr))
	fmt.Println(reverse2(arr))

	fmt.Println("Merge Sort")
	fmt.Println(mergeSort([]int{0, 3, 4, 31}, []int{4, 6, 10, 20}))

	fmt.Println("Sun Maximus")
	fmt.Println(sunMaximus([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
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

func mergeSort(arrayOne []int, arrayTwo []int) []int {
	var result []int
	i := 0
	j := 0
	for i < len(arrayOne) && j < len(arrayTwo) {
		if arrayTwo[j] > arrayOne[i] {
			result = append(result, arrayOne[i])
			i++
		} else {
			result = append(result, arrayTwo[j])
			j++
		}
	}
	for i < len(arrayOne) {
		result = append(result, arrayOne[i])
		i++
	}
	for j < len(arrayTwo) {
		result = append(result, arrayTwo[j])
		j++
	}
	return result
}

func sunMaximus(nums []int) int {
	localSum, globalSum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if localSum < 0 {
			localSum = 0
		}
		localSum += nums[i]

		if globalSum < localSum {
			globalSum = localSum
		}
	}

	return globalSum
}
