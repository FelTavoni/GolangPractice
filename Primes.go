/*
 * Created by Felipe Tavoni on Mon May 10 2021
 *
 * Problem proposal extracted from LeetCode©
 *
 * Count Primes
 * Count the number of prime numbers less than a non-negative number, n.
 *
 * My solution: As previous studied by me in Programação Paralela Distribuída (Parallel Computing), available in my github,
 * 	the sieve of erastothenes is a efficienty approach for this problem, because it doesn't iterate over all possible primes.
 * Find more at my study on my githuhub!
 */

package main

import (
	"fmt"
)

func countPrimes(n int) int {

	if n == 0 || n == 1 {
		return 0
	}

	// Use an array to process all non-prime number.
	var numbers = make([]int, n) // Dynamically allocating an array in Go. Don't include n.
	var count int                // Holds the number of primes in the array.

	// Filling the array with number from 0 to n.
	for i := 2; i < n; i++ {
		numbers[i] = i
	}

	// So, by using the Erastothenes algorithm, process the 'numbers' array, removing all non-prime elements.
	for i := 2; i < n; i++ {
		if numbers[i] != 0 { // This number is a prime!
			count++
			for j := i * i; j < n; j += i {
				if j%i == 0 { // Remove all multiples of this number
					numbers[j] = 0
				}
			}
		}
	}

	fmt.Println("Primes are: ", numbers)

	return count
}

func main() {
	var n int
	fmt.Print("Should I check all the prime-numbers until...: ")
	fmt.Scan(&n)
	fmt.Println("There are", countPrimes(n), "primes from 0 to", n)
}
