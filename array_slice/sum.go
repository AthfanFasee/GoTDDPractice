package arrayslice

func Sum(slice []int) int {
	var total int

	for _, number := range slice {
		total += number
	}

	return total
}

func SumAll(NumbersToSum ...[]int) []int {
	sumAll := make([]int, len(NumbersToSum))

	for i, slice := range NumbersToSum {
		sumAll[i] = Sum(slice)
	}

	return sumAll
}

// I think version Two takes more time, bcs of the need for reallocation.
func SumAllVersinTwo(NumbersToSum ...[]int) []int {
	var sumAll []int

	for _, slice := range NumbersToSum {
		sumAll = append(sumAll, Sum(slice))
	}

	return sumAll
}

// SumAllTails will calculate the totals of the "tails" of each slice.
// The tail of a collection is all items in the collection except the first one (the "head")
func SumAllTails(NumbersToSum ...[]int) []int {
	var sumAll []int

	for _, slice := range NumbersToSum {
		if len(slice) == 0 {
			sumAll = append(sumAll, 0)
		} else {
			sumAll = append(sumAll, Sum(slice[1:]))
		}
	}

	return sumAll
}