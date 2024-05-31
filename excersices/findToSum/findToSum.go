package findToSum

func TwoSum(nums []int, sum int) bool {
	dictComplement := make(map[int]bool)
	for _, num := range nums {
		complement := sum - num
		_, ok := dictComplement[complement]
		if ok {
			return true
		} else {
			dictComplement[num] = true
		}

	}
	return false
}
