package arraysandslices

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numSets ...[]int) (sumSet []int) {
	for _, numSet := range numSets {
		sumSet = append(sumSet, Sum(numSet))
	}
	return
}
