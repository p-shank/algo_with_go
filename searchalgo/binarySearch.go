package searchalgo

import "fmt"

// func main() {
// 	items := []int{1, 2, 3, 4, 5, 6, 7}
// 	position, steps := search(items, 5)
// 	fmt.Println("position of element", position)
// 	fmt.Println("total steps: ", steps)
// }

func search(items []int, item int) (int, int) {
	low, high := 0, len(items)-1
	var count int = 0
	fmt.Println("start -> ", items)
	for low <= high {
		count += 1
		var mid int = (low + high) / 2
		if items[mid] == item {
			return mid, count
		} else if items[mid] < item {
			low = mid + 1
		} else {
			high = mid - 1
		}
		fmt.Printf("iteration %d -> %d \n", count, items[low:high+1])
	}
	return -1, count
}
