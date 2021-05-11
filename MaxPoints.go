/*
 * Created by Felipe Tavoni on Mon May 10 2021
 *
 * Problem proposal extracted from LeetCode©
 *
 * Maximum Points You Can Obtain From Cards
 *
 * There are several cards arranged in a row, and each card has an associated number of points The points are given in the integer array cardPoints.
 * In one step, you can take one card from the beginning or from the end of the row. You have to take exactly k cards.
 * Your score is the sum of the points of the cards you have taken.
 * Given the integer array cardPoints and the integer k, return the maximum score you can obtain.
 *
 * My solution: In this approach, if we imagine a floating window that moves over the array, we can compute all possibles solution.
 *	It'd be something like: [0 1 2) 3 4 5()] -> [0 1) 2 3 4 (5] -> [0) 1 2 3 (4 5] -> [)0 1 2 (3 4 5]
 */

package main

import (
	"fmt"
)

func maxScore(cardPoints []int, k int) int {
	// Quantity of lements in the array.
	var n int = len(cardPoints)
	// The total pontuation on a iteration, the best pontuation among all others, and a auxiliar to compute the rightmost sum.
	var total, best, aux int

	// Get a initial total points possible.
	for i := 0; i < k; i++ {
		total += cardPoints[i]
	}
	best = total

	// Moving the window, we subtract the k - i element from the left and add the n - k + i element on the right.
	for i := k - 1; i >= 0; i-- {
		total -= cardPoints[i]
		aux += cardPoints[n-k+i]
		if total+aux > best {
			best = total + aux
		}
	}

	return best
}

func main() {
	var n, k int
	fmt.Print("How many cards? ")
	fmt.Scan(&n)
	var cardPoints = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Print("cardPoints[", i, "]:")
		fmt.Scan(&cardPoints[i])
	}
	fmt.Print("The number of cards you can take: ")
	fmt.Scan(&k)
	fmt.Println(maxScore(cardPoints, k))
}
