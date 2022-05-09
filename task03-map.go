package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	var keys []int = make([]int, 0, len(input))
	for k := range input {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	var sorted_values []string = make([]string, 0, len(input))
	for elem := range keys {
		sorted_values = append(sorted_values, input[elem])
	}

	return sorted_values
}
