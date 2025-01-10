package servicepackage

const repeatCount = 5

func BenchmarkRepeat() string {
	data := Repeat("a")
	return data
}

func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}
