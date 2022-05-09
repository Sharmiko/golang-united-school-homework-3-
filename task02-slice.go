package homework

import "sort"

func reverse(input []int64) (result []int64) {
	sort.Slice(input, func(x, y int) bool {
		return input[x] < input[y]
	})
	return input
}
