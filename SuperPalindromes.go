/*
 * Created by Felipe Tavoni on Sat May 08 2021
 *
 * Problem proposal extracted from LeetCode©
 *
 * Super Palindromes
 * Let's say a positive integer is a super-palindrome if it is a palindrome, and it is also the square of a palindrome.
 * Given two positive integers left and right represented as strings, return the number of super-palindromes integers
 *	in the inclusive range [left, right].
 *
 * My solution: As a previous exercise, find the palindrome of the number, check if is in the especified range, get this
 *  number square root, and check if it's a palindrome. This algorithm has a efficiency of time O(n), and a space efficiency
 *	of O(1). Not too efficient for leetcode...
 */

package main

import (
	"fmt"
	"math"
	"strconv"
)

// Auxiliar function to find out if the string is a palindrome. Compares rune by rune.
func checkPalindrome(num string) bool {
	var i, j int
	j = len(num) - 1
	for i < len(num) {
		if num[i] != num[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func superpalindromesInRange(left string, right string) int {
	// Left and right boundaries, along with a 'number' which will contain the square root to be evaluated.
	var leftInt, rightInt, number, i int64
	// Error variable just for read.
	var _ error
	// The total of super-palindromes in this range.
	var total int
	// Uncomment to know which numbers are super-palindromes.
	// var numbers []int64

	// Convert string boundaries to int.
	leftInt, _ = strconv.ParseInt(left, 10, 0)
	rightInt, _ = strconv.ParseInt(right, 10, 0)

	// Check range palindromes
	for i = 0; i <= int64(math.Sqrt(float64(rightInt))); i++ {
		// Convert the number to a string.
		numString := strconv.FormatInt(i, 10)
		// Check if it's a palindrome.
		if checkPalindrome(numString) {
			// Get the square of the palindrome.
			number = i * i
			// As the integer truncates the number at the conversion, we need to compare with it's new square, that might be different.
			if leftInt <= number && number <= rightInt {
				numString = strconv.FormatInt(number, 10)
				// Then, check again if the square root is a palindrome.
				if checkPalindrome(numString) {
					// Uncomment to know which numbers are super-palindromes.
					// numbers = append(numbers, number)
					// If it's, count it.
					total++
				}
			}
		}
	}

	// Uncomment to know which numbers are super-palindromes.
	// fmt.Println("numbers = ", numbers)

	return total
}

func main() {
	var left, right string
	fmt.Print("Enter the left position of the range:")
	fmt.Scan(&left)
	fmt.Print("Enter the right position of the range:")
	fmt.Scan(&right)
	fmt.Println(superpalindromesInRange(left, right))
}
