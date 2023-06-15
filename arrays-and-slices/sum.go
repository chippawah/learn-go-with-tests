package arraysandslices

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numSets ...[]int) (sumSet []int) {
	numSetLen := len(numSets)
	sumSet = make([]int, numSetLen)
	for idx, numSet := range numSets {
		sumSet[idx] = Sum(numSet)
	}
	return
}
