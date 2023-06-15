package arraysandslices

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAllTails(numSets ...[]int) (tailSumSet []int) {
	for _, numSet := range numSets {
		if len(numSet) == 0 {
			numSet = []int{0}
		}
		tail := numSet[1:]
		tailSumSet = append(tailSumSet, Sum(tail))
	}
	return
}
