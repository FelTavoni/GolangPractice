/*
 * Created by Felipe Tavoni on Thu May 13 2021
 *
 * Problem proposal extracted from LeetCode©
 *
 * Jewels And Stones
 *
 * You're given strings jewels representing the types of stones that are jewels, and stones representing the stones you have.
 * Each character in stones is a type of stone you have. You want to know how many of the stones you have are also jewels.
 * Letters are case sensitive, so "a" is considered a different type of stone from "A".
 *
 * My solution: Basic solution. For each stone, check if we have a matching jewel.
 */

package main

import (
	"fmt"
)

func numJewelsInStones(jewels string, stones string) int {
	var count int

	for _, stone := range stones {
		for _, jewel := range jewels {
			if jewel == stone {
				count++
			}
		}
	}

	return count
}

func main() {
	var jewels, stones string
	fmt.Print("Jewels = ")
	fmt.Scan(&jewels)
	fmt.Print("Stones = ")
	fmt.Scan(&stones)
	fmt.Println(numJewelsInStones(jewels, stones))
}
