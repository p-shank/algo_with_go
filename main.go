package main

import (
	"fmt"
	"sortalgo"
)

func main() {
	items := []int{2, 3, 5, 1, 4}
	sorted := sortalgo.Sort(items)

	fmt.Println("sorted array: ", sorted)
}
