/*
 * Created by Felipe Tavoni on Sat May 01 2021
 *
 * Problem proposal extracted from LeetCode©
 *
 * Reverse Integer
 * Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the
 *	signed 32-bit integer range [-231, 231 - 1], then return 0.
 *
 * My solution: Mod x to obtain the least number, while multiplying the result by ten, so we can move the least number to
 * 	it's maximum position.
 */

package main

import (
	"fmt"
)

func reverse(x int) int {
	var maximum int64 = 2147483647
	var num int
	var negative bool

	if x < 0 {
		x = x * (-1)
		negative = true
	}

	for x > 0 {
		num = (num * 10) + (x % 10)
		x /= 10
	}

	if int64(num) > maximum {
		return 0
	}

	if negative {
		num = num * (-1)
	}

	return num
}

func main() {
	var number int
	fmt.Print("Type the number to be reversed: ")
	fmt.Scan(&number)
	fmt.Print(reverse(number))
}
