/*
 * Created by Felipe Tavoni on Mon May 03 2021
 *
 * Problem proposal extracted from LeetCode©
 *
 * Running Sum of 1d Array
 * Given an array nums. We define a running sum of an array as runningSum[i] = sum(nums[0]…nums[i]).
 * Return the running sum of nums.
 *
 * My Solution: Like a fibonacci algorithm, sum with the previous element. We can do this in a O(n) efficiency.
 *
 */

package main

import (
	"fmt"
)

// Checks if a number n is a power of three
func runningSum(nums []int) []int {
	for i, _ := range nums {
		if i+1 < len(nums) {
			nums[i+1] += nums[i]
		}
	}
	return nums
}

func main() {
	var i, size int
	fmt.Println("How many numbers will your array possess?")
	fmt.Scan(&size)
	var nums = make([]int, size)
	for i < size {
		fmt.Print("nums[", i, "]:")
		fmt.Scan(&nums[i])
		i++
	}
	fmt.Println(runningSum(nums))
}
