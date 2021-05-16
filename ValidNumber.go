/*
 * Created by Felipe Tavoni on Sat May 15 2021
 *
 * Problem proposal extracted from LeetCode©
 *
 * A valid number can be split up into these components (in order):
 *	1. A decimal number or an integer.
 *	2. (Optional) An 'e' or 'E', followed by an integer.
 *
 * A decimal number can be split up into these components (in order):
 *	1. (Optional) A sign character (either '+' or '-').
 *	2. One of the following formats:
 *		a. At least one digit, followed by a dot '.'.
 *		b. At least one digit, followed by a dot '.', followed by at least one digit.
 * 		c. A dot '.', followed by at least one digit.
 *
 * An integer can be split up into these components (in order):
 * 	1. (Optional) A sign character (either '+' or '-').
 *  2. At least one digit.
 *
 * Given a string s, return true if s is a valid number.
 * Example: "2e10"
 *
 * My Solution: Solution a little simple. For each caracter read, check all rules of the integers and decimals. We can do that
 *	by using some boolean variables. This solution is O(n).
 *
 */

package main

import (
	"fmt"
	"unicode"
)

func isNumber(s string) bool {

	var num, exp, sign, dec bool

	for _, n := range s {
		// If it's a digit, ok...(the 'char' in Go is a rune, represents Unicode).
		if unicode.IsDigit(n) {
			num = true
			// If we found a dot...
		} else if n == '.' {
			// Must be a unique dot or before the exponential symbol.
			if dec || exp {
				return false
				// Otherwise, mark it as seen.
			} else {
				dec = true
			}
			// If founded a exponential...
		} else if n == 'e' || n == 'E' {
			// It has to be also unique and verify if there is not a number befor the exponential.
			if exp || !num {
				return false
			} else {
				// Reseting a few variables so it's possible to analyse the number after the exponential.
				exp = true
				num = false
				sign = false
				dec = false
			}
			// Found a signal...
		} else if n == '+' || n == '-' {
			// Has to be unique
			if sign || num || dec {
				return false
			} else {
				sign = true
			}
			// None of above, for example a letter...
		} else {
			return false
		}
	}

	return num
}

func main() {
	var number string
	fmt.Print("Enter the number to be analysed: ")
	fmt.Scan(&number)
	if isNumber(number) {
		fmt.Println("It is a valid number!")
	} else {
		fmt.Println("It isn't a valid number...")
	}
}
