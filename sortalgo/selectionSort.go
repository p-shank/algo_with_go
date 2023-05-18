package sortalgo

func Sort(items []int) []int {

	for i := 0; i < len(items); i++ {
		idx := i
		for j := i + 1; j < len(items); j++ {
			if items[idx] >= items[j] {
				idx = j
			}
		}
		tempItem := items[i]
		items[i] = items[idx]
		items[idx] = tempItem
	}

	return items
}
