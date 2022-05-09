package homework

func reverse(input []int64) (result []int64) {
	var arrSize int = len(input) - 1
	for i := 0; i < arrSize; i = i + 1 {
		input[i], input[arrSize] = input[arrSize], input[i]
		arrSize = arrSize - 1
	}

	return input
}
