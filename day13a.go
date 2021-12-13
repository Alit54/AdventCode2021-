package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var matrix [1400][1400]bool
	var x, y int
	var req string
	// Read the input and set the position to TRUE
	file, err := os.Open("day 13 input.txt")
	if err != nil {
		panic("You need to create a file named 'day 13 input.txt' in order to continue!")
	}
	defer file.Close()
	for {
		_, err = fmt.Fscanf(file, "%d", &x)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Fscanf(file, "%d", &y)
		matrix[y][x] = true
		fmt.Fscanf(file, "%s")
	}
	// Read the first instruction: You need to create a symmetric ax based on what you read
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	req = scanner.Text()
	fmt.Println(req)
}
