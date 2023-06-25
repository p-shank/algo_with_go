package sort_algo

import "fmt"

func Sort(items []int) []int {
	count := 0
	for i := 0; i < len(items); i++ {
		idx := i
		count = count + 1
		for j := i + 1; j < len(items); j++ {
			count = count + 1
			if items[idx] >= items[j] {
				idx = j
			}
		}
		tempItem := items[i]
		items[i] = items[idx]
		items[idx] = tempItem
	}

	fmt.Println("no of operations: ", count)
	return items
}
