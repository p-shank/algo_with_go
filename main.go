package main

import (
	"algo_with_go/sortalgo"
	"fmt"
)

func main() {
	items := []int{5, 4, 3, 2, 1}
	sorted := sortalgo.Sort(items)
	fmt.Println("sorted array: ", sorted)
}
