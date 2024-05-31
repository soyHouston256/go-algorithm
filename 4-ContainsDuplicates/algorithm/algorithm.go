package algorithm

func ContainsDuplicate(nums []int) bool {
	maps := make(map[int]bool)
	for _, num := range nums {
		if maps[num] {
			return true
		}
		maps[num] = true
	}
	return false
}
