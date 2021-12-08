package main

import (
	"fmt"
)

func main() {
	var n uint8
	var lanternfish []uint8
	// Memorise the lanternfishes
	for {
		_, err := fmt.Scan(&n)
		if err != nil {
			break
		}
		lanternfish = append(lanternfish, n)
	}
	// repeat until you reach day 80
	for days := 1; days <= 80; days++ {
		for i := 0; i < len(lanternfish); i++ {
			if lanternfish[i] == 0 {
				lanternfish[i] = 7
				// check if some index reaches 0. If so, append a new blowfish with value 8
				lanternfish = append(lanternfish, 9)
			}
		}
		// decrease each value
		for i := 0; i < len(lanternfish); i++ {
			lanternfish[i]--
		}
	}
	// Your answer is the lenght of the slice
	fmt.Println(len(lanternfish))
}
