package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		input        string
		winnerScore  int
		winnerNumber int
	)

	reader := bufio.NewReader(os.Stdin)
	for i := 1; i <= 5; i++ {
		input, _ = reader.ReadString('\n')

		numbs := strings.Fields(input)
		tempScore := 0
		for _, v := range numbs {
			i, _ := strconv.Atoi(v)
			tempScore += i
		}

		if tempScore > winnerScore {
			winnerScore = tempScore
			winnerNumber = i
		}
	}

	fmt.Printf("%d %d", winnerNumber, winnerScore)
}
