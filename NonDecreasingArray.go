/*
 * Created by Felipe Tavoni on Tue May 04 2021
 *
 * Problem proposal extracted from LeetCode©
 *
 * Non-decreasing Array
 * Given an array nums with n integers, your task is to check if it could become non-decreasing by modifying at most one element.
 * We define an array is non-decreasing if nums[i] <= nums[i + 1] holds for every i (0-based) such that (0 <= i <= n - 2).
 *
 * My Solution: The problem can be solved with a O(n) efficiency by checking in runtime the next and previous elements and change
 *	if necessary to avoid future impacts.
 *
 */

package main

import (
	"fmt"
)

func checkPossibility(nums []int) bool {
	var changes int
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			if i > 0 {
				if nums[i-1] <= nums[i+1] {
					nums[i] = nums[i-1]
				} else {
					nums[i+1] = nums[i]
				}
			}
			changes++
			if changes > 1 {
				return false
			}
		}
	}
	return true
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
	fmt.Println(checkPossibility(nums))
}
