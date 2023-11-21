package services

func FindOdd(input []int) int {
	result := 0
	for _, num := range input {
		result ^= num
	}

	return result
}
