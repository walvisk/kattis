package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		x int
		n int
		r int
	)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "%d", &x)

		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "%d", &n)

		for i := 0; i < n; i++ {
			var ti, tx int
			scanner.Scan()
			fmt.Sscanf(scanner.Text(), "%d", &ti)

			tx = x
			r = r + tx - ti
		}
		break
	}

	r = r + x
	fmt.Print(r)
}
