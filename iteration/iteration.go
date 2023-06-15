package iteration

func Repeat(char string, count int) (repeated string) {
	if count <= 1 {
		count = 1
	}
	for i := 0; i < count; i++ {
		repeated += char
	}
	return
}
