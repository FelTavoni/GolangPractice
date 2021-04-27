/*
 * Created by Felipe Tavoni on Tue Apr 27 2021
 *
 * Problem proposal extracted from LeetCode©
 *
 * Power of Three
 * Given an integer n, return true if it is a power of three. Otherwise, return false.
 * An integer n is a power of three, if there exists an integer x such that n == 3x.
 */

package main

import (
	"fmt"
)

// Checks if a number n is a power of three
// My solution
func isPowerOfThree(n int) bool {
	// Avoids inf loop
	if n == 0 {
		return false
	}
	for n%3 == 0 {
		n = n / 3
	}

	return n == 1
}

// Checks if a number n is a power of three
// The max possible number power of 3 is 1162261467. So basically divide it by n to check if n is also multiple of 3.
func isPowerOfThreeFaster(n int) bool {
	return n > 0 && 1162261467%n == 0
}

func main() {
	var number int
	fmt.Print("Which number should I check if it's a power of 3? ")
	fmt.Scan(&number)
	if isPowerOfThree(number) {
		fmt.Println("It's a power of three!")
	} else {
		fmt.Print("Sorry...it isn't a power of three...")
	}
}
