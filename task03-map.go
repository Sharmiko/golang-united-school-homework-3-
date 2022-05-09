package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	var keys []int = make([]int, 0, len(input))
	for k := range input {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	var sortedValues []string = make([]string, 0, len(input))
	for elem := range keys {
		sortedValues = append(sortedValues, input[elem])
	}

	return sortedValues
}
