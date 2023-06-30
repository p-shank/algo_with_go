package sort_algo

import "fmt"

func Qsort(items []int) []int {

	if len(items) <= 1 {
		return items
	} else {
		pivot := items[len(items)/2]
		var lesser []int
		var greater []int

		n := len(items)

		for i := 0; i < n; i++ {
			if items[i] < pivot {
				lesser = append(lesser, items[i])
			}
		}

		for i := 0; i < n; i++ {
			if items[i] > pivot {
				greater = append(greater, items[i])
			}
		}

		fmt.Println(lesser, pivot, greater)

		return append(append(Qsort(lesser), pivot), Qsort(greater)...)

	}
}
