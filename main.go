package main

import "github.com/RohBot/GOLANG/GoPhase3/leetcode"

func main() {
	leetcode.RemoveDuplicatesFromSortedArray([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
}

// Input: nums = [0,0,1,1,1,2,2,3,3,4] 0,1,2,3,4,2,2,3,3
// Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
