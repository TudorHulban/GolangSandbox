package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/two-sum/

// Given an array of integers nums and an integer target,
// return indices of the two numbers such that they add up to target.

// Example 1:
// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Output: Because nums[0] + nums[1] == 9, we return [0, 1].

// Example 2:
// Input: nums = [3,2,4], target = 6
// Output: [1,2]

// Example 3:
// Input: nums = [3,3], target = 6
// Output: [0,1]

// Constraints:

// 2 <= nums.length <= 103
// -109 <= nums[i] <= 109
// -109 <= target <= 109
// Only one valid answer exists.

func main() {
	a := []int{1, 2, 3}
	t := 5

	fmt.Println(indexes(a, t))
}

// indexes Returns the single pair matching target. No error handling.
func indexes(a []int, target int) [2]int {
	sort.Ints(a)

	for i := range a {
		for j := i + 1; j < len(a) && a[i]+a[j] <= target; j++ {
			if a[i]+a[j] == target {
				return [2]int{i, j}
			}
		}
	}

	return [2]int{}
}
