package divide_and_conquer

func MaxElement(item []int) int {

	if len(item) == 1 {
		return item[0]
	}

	recItem := MaxElement(item[1:])

	if item[0] > recItem {
		return item[0]
	}

	return recItem

}
