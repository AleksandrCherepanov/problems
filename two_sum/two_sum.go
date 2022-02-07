package two_sum

import (
	"sort"
)

func Execute(numbers []int, target int) [2]int {
	sort.Ints(numbers)

	i := 0
	j := len(numbers) - 1

	result := [2]int{}

	for i < j {
		sum := numbers[i] + numbers[j]
		if sum == target {
			result[0] = i
			result[1] = j
			break
		}

		if sum > target {
			j--
		} else {
			i++
		}
	}

	return result
}
