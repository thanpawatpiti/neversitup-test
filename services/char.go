package services

func SwitchChar(s string, i int, result *[]string) {
	if i == len(s) {
		*result = append(*result, s)
		return
	}

	for j := i; j < len(s); j++ {
		s = swap(s, i, j)
		SwitchChar(s, i+1, result)
		s = swap(s, i, j)
	}
}

func swap(s string, i, j int) string {
	b := []byte(s)
	b[i], b[j] = b[j], b[i]
	return string(b)
}
