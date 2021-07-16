/*
 * Created by Felipe Tavoni on Fri Jul 16 2021
 *
 * Problem proposal extracted from LeetCode©
 *
 * 4Sum
 *
 * Given an array nums of n integers, return an array of all the unique quadruplets [nums[a], nums[b], nums[c], nums[d]] such that:
 * 	0 <= a, b, c, d < n
 * 	a, b, c, and d are distinct.
 *	nums[a] + nums[b] + nums[c] + nums[d] == target
 * You may return the answer in any order.
 *
 * My solution: If we fix 2 numbers, we only need to find the other two numbers of the solution. It's a brute force solution. For the
 *	duplicate numbers, we simply iterate over them to not obtain more than one equal solution.
 */

package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	ans := [][]int{}
	sort.Ints(nums)
	var size int = len(nums)
	var i, j int = 0, 0
	// Fixing the first number.
	for i < size {
		j = i + 1
		// Fixing the second number
		for j < size {
			var start, finish, sum int = j + 1, size - 1, target - nums[i] - nums[j]
			// For the third and fourth number, inside the space we've left, search for the solution.
			for start < finish {
				if nums[start]+nums[finish] == sum {
					ans = append(ans, []int{nums[i], nums[j], nums[start], nums[finish]})
					var third, fourth int = nums[start], nums[finish]
					for start < finish && nums[start] == third {
						start++
					}
					for start < finish && nums[finish] == fourth {
						finish--
					}
				} else if nums[start]+nums[finish] > sum {
					finish--
				} else {
					start++
				}
			}
			// Skipping duplicates.
			for j+1 < size && nums[j+1] == nums[j] {
				j++
			}
			j++
		}
		// Also skipping duplicates.
		for i+1 < size && nums[i+1] == nums[i] {
			i++
		}
		i++
	}
	return ans
}

func main() {
	var i, size, target int
	fmt.Print("How many numbers will your array possess? ")
	fmt.Scan(&size)
	var nums = make([]int, size)
	for i < size {
		fmt.Print("nums[", i, "]:")
		fmt.Scan(&nums[i])
		i++
	}
	fmt.Print("Target = ")
	fmt.Scan(&target)
	fmt.Println("nums = ", nums, "& target =", target)
	fmt.Println(fourSum(nums, target))
}
