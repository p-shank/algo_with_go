package divide_and_conquer

func SumArrayItems(item []int) int {
	if len(item) == 0 {
		return 0
	}
	return item[0] + SumArrayItems(item[1:])
}
