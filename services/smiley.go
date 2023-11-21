package services

import "regexp"

func CountSmileys(arr []string) int {
	regex := regexp.MustCompile(`^[:;][-~]?[)D]$`)

	count := 0
	for _, face := range arr {
		if regex.MatchString(face) {
			count++
		}
	}
	return count
}
