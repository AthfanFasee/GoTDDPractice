package iteration

func Repeat(x string, y int) string {
	var repeatedString string

	for i := 0; i < y; i++ {
		repeatedString += x
	}

	return repeatedString
}