package homework

import "sort"

func reverse(input []int64) (result []int64) {
	isSorted := sort.SliceIsSorted(input, func(x, y int) bool {
		return input[x] < input[y]
	})
	if isSorted {
		sort.Slice(input, func(x, y int) bool {
			return input[x] > input[y]
		})
	} else {
		sort.Slice(input, func(x, y int) bool {
			return input[x] < input[y]
		})
	}

	return input
}
