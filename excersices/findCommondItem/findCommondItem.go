package findCommondItem

func FindCommonItem(array1 []string, array2 []string) bool {
	var mapItems = make(map[string]bool)
	for _, v := range array1 {
		if _, ok := mapItems[v]; !ok {
			mapItems[v] = true
		}

	}

	for _, v := range array2 {
		_, ok := mapItems[v]
		if ok {
			return true
		}
	}
	return false
}
