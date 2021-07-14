/*
 * Created by Felipe Tavoni on Wed Jul 14 2021
 *
 * Problem proposal extracted from LeetCode©
 *
 * Custom Sort String
 *
 * order and str are strings composed of lowercase letters. In order, no letter occurs more than once.
 * order was sorted in some custom order previously. We want to permute the characters of str so that they match the order that order 
 *	was sorted. More specifically, if x occurs before y in order, then x should occur before y in the returned string.
 * Return any permutation of str (as a string) that satisfies this property.
 *
 * My solution: Basically, we can count the occurence of each letter by using a hash, and then simply rebuild the string following 
 *	the order, adding the unmapped letters at the end.
 */

package main

import (
	"fmt"
)

func customSortString(order string, str string) string {
	// Creates a hash that holds the occurence of each letter (in Go, rune) in 'str'.
	hash := map[rune]int{}
	for _, let := range str {
		hash[let]++
	}

	// Now, we start to rebuilt the return expression checking the 'order' string.
	// Also delete some string so that the remaining ones are the letters that are not included in 'order'.
	var retStr string
	var qnt int
	for _, let := range order {
		qnt = hash[let]
		for i := 0; i < qnt; i++ {
			retStr = retStr + string(let)
			delete(hash, let)
		}
	}

	// At the end, simply append the remaining letters.
	for let, qnt := range hash {
		fmt.Println(let)
		for i := 0; i < qnt; i++ {
			retStr = retStr + string(let)
		}
	}

	// Returning the expression.
	return retStr
}

func main() {
	var order, str string
	fmt.Print("Order: ")
	fmt.Scan(&order)
	fmt.Print("Str: ")
	fmt.Scan(&str)
	fmt.Println(customSortString(order, str))
}
