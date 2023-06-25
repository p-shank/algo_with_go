package main

import (
	"algo_with_go/divide_and_conquer"
	"fmt"
)

// func main() {
// 	items := []int{5, 4, 3, 2, 1}
// 	sorted := sort_algo.Sort(items)
// 	fmt.Println("sorted array: ", sorted)
// }

func main() {
	items := []int{1, 4, 3, 5, 2, 7}
	sum := divide_and_conquer.SumArrayItems(items)
	count := divide_and_conquer.NoOfElement(items)
	max := divide_and_conquer.MaxElement(items)
	fmt.Println("sumed up array: ", sum)
	fmt.Println("no of element: ", count)
	fmt.Println("max element: ", max)
}
