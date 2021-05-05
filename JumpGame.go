/*
 * Created by Felipe Tavoni on Wed May 05 2021
 *
 * Problem proposal extracted from LeetCode©
 *
 * Jump Game II
 * Given an array of non-negative integers nums, you are initially positioned at the first index of the array.
 * Each element in the array represents your maximum jump length at that position.
 * Your goal is to reach the last index in the minimum number of jumps.
 * You can assume that you can always reach the last index.
 *
 * My Solution: A greedy aproach is capable to solve this problem with a O(n) efficiency. While it iterates
 *	over the array, it keeps tracking the highest possible jump by using math.max.
 *
 */

package main

import (
	"fmt"
	"math"
)

func jump(nums []int) int {
	// If there's no place to jump, or not enough jumps, return.
	if len(nums) == 1 || nums[0] == 0 {
		return 0
	}
	// Otherwise, there's at least 1 jumo to be made
	var curr, highest, jumps int
	for i := 0; i < len(nums); i++ {
		// Avoiding "out of index"
		if i == len(nums)-1 {
			return jumps
		}
		// Updating the maximum place we can jump
		highest = int(math.Max(float64(highest), float64(i+nums[i])))
		// Counting the number of jumps
		if i == curr {
			jumps++
			curr = highest
		}
	}
	return jumps
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
	fmt.Println(jump(nums))
}
