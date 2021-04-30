/*
 * Created by Felipe Tavoni on Fri Apr 30 2021
 *
 * Problem proposal extracted from LeetCode©
 *
 * Powerful Integers
 * Given three integers x, y, and bound, return a list of all the powerful integers that have a value less than or equal to bound.
 * An integer is powerful if it can be represented as x^i + y^j for some integers i >= 0 and j >= 0.
 * You may return the answer in any order. In your answer, each value should occur at most once.
 *
 * My Solution: The only possible approach for this involves brute force.
 *
 */

package main

import (
	"fmt"
	"math"
)

// Returns the powerful integers contained in the bound.
func powerfulIntegers(x int, y int, bound int) []int {
	// Converting the arguments into float.
	var floatX, floatY, floatBound float64 = float64(x), float64(y), float64(bound)
	// Declaring a slice that hols the founded powerful integers.
	var integers []int = make([]int, 0)
	// A hashmap, that help avoids adding twice a number.
	var integersBool = make(map[float64]bool)
	// The calculated number.
	var num float64

	// Brute Force...
	for i := 0.0; math.Pow(floatX, i) <= floatBound; i++ {
		for j := 0.0; math.Pow(floatY, j) <= floatBound; j++ {
			num = math.Pow(floatX, i) + math.Pow(floatY, j)
			// If the number is between 0 - 'bound' and has not been added, add it!
			if num <= floatBound && !integersBool[num] {
				integers, integersBool[num] = append(integers, int(num)), true
			}
			// No point in considering other powers of "1".
			if y == 1 {
				break
			}
		}
		// Same for x...
		if x == 1 {
			break
		}
	}

	// Return the powerful integers!
	return integers
}

func main() {
	var x, y, bound int
	fmt.Print("Value of x? ")
	fmt.Scan(&x)
	fmt.Print("Value of y? ")
	fmt.Scan(&y)
	fmt.Print("Value of bound? ")
	fmt.Scan(&bound)
	fmt.Printf("The powerful integers in %d are:", bound)
	fmt.Print(powerfulIntegers(x, y, bound))
}
