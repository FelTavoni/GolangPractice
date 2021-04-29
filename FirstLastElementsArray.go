/*
 * Created by Felipe Tavoni on Thu Apr 29 2021
 *
 * Problem proposal extracted from LeetCode©
 *
 * Find first and last position of elements in sorted array
 * Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value.
 * If target is not found in the array, return [-1, -1].
 *
 * My solution: Basically, run a binary search - since the array is sorted - to find the target number position. After
 *	that, check for previous indexes, because the index returned by the binary search may not be the first. Then, search
 *	for the last and return the indexes. This algorithm has a O(log n) efficiency, provided by the binary search!
 */

package main

import (
	"fmt"
)

func binarySearch(arr []int, l int, r int, x int) int {
	if r >= l {
		var mid int = l + (r-l)/2

		// If the element is present at the middle itself
		if arr[mid] == x {
			return mid
		}

		// If element is smaller than mid, then it can only be present in left subarray
		if arr[mid] > x {
			return binarySearch(arr, l, mid-1, x)
		}

		// Else the element can only be present in right subarray
		return binarySearch(arr, mid+1, r, x)
	}

	// We reach here when element is not present in array
	return -1
}

func searchRange(nums []int, target int) []int {

	// Call a binary search to search for the target number.
	var i int = binarySearch(nums, 0, len(nums)-1, target)

	// If the number doesn't exists...return [-1, -1]
	if i == -1 {
		return []int{-1, -1}
	}

	var first, last int

	// The target number found in the binary search, if repeated in the array, may not be the first...so we check
	//	for previous numbers if there's any...
	for i > 0 && nums[i-1] == target {
		i--
	}

	// Get it's position
	first = i

	// Search the last element
	for i < len(nums)-1 && nums[i+1] == target {
		i++
	}

	// Marks it's position
	last = i

	// Returns the indexes
	return []int{first, last}
}

func main() {
	var i, size, target int
	// nums := []int
	fmt.Println("How many numbers will your array possess?")
	fmt.Scan(&size)
	var nums = make([]int, size)
	for i < size {
		fmt.Print("nums[", i, "]:")
		fmt.Scan(&nums[i])
		i++
	}
	fmt.Println("Typed array:", nums)
	fmt.Println("Which number should I search in this array? ")
	fmt.Scan(&target)
	fmt.Println(searchRange(nums, target))
}
