package main

import "fmt"

func main() {
	var (
		input  string
		output []byte
	)

	fmt.Scanln(&input)
	output = append(output, input[0])
	for i := 1; i < len(input); i++ {
		if input[i] != input[i-1] {
			output = append(output, input[i])
		}
	}

	fmt.Print(string(output))
}
