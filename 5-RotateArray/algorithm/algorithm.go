package algorithm

import "fmt"

func Rotate(nums []int, k int) {
	if len(nums) == 1 || nums == nil {
		return
	}
	k = k % len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}
func reverse(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func Rotate2(nums []int, k int) {
	fmt.Println(nums)
	for k > 0 {
		k = k % len(nums)
		temp := nums[0]
		nums[0] = nums[len(nums)-1]
		for i := 1; i < len(nums); i++ {
			temp2 := nums[i]
			nums[i] = temp
			temp = temp2
		}
		k--
		fmt.Println(nums)
	}
}
