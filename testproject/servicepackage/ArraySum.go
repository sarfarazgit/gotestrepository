package servicepackage

func SumofArrayElements(elements []int) int {
	sum := 0
	for _, number := range elements {
		sum += number
	}
	return sum
}
