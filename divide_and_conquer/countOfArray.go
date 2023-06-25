package divide_and_conquer

func NoOfElement(item []int) int {
	if len(item) == 0 {
		return 0
	}
	return 1 + NoOfElement(item[1:])

}
