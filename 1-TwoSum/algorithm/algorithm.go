package algorithm

func TwoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		_, ok := numMap[complement]
		if ok {
			return []int{numMap[complement], i}
		}
		numMap[num] = i
	}
	return []int{0, 0}
}
