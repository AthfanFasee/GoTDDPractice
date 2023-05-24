package arrayslice

func Sum(slice []int) int {
	var total int

	for _, number := range slice {
		total += number
	}

	return total
}