package main

import (
	"fmt"
	"gotest/services"
)

func main() {
	var input string
	fmt.Print("Enter a string: ")
	fmt.Scanf("%s", &input)

	result := []string{}
	if len(input) > 0 {
		services.SwitchChar(input, 0, &result)
	}

	fmt.Println("Result test 1: ", result)

	input2 := []int{1, 1, 2, 3, 3, 3, 4, 4, 5}
	result2 := services.FindOdd(input2)

	fmt.Println("Result	test 2: ", result2)

	fmt.Println(services.CountSmileys([]string{":)", ";D", ":~)", ";-D"}))
}
