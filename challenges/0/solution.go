package challenges

import (
	"sort"
)

// Solution given an array A of N integers, returns the smallest positive integer (greater than 0) that does not occur in A
func Solution(A []int) int {
	sort.Ints(A) // the most expensive part of this entire process

	length := len(A)
	answer := A[length-1] + 1 // prime the answer, regardless if it's the final answer

	// end quickly on the most obvious condition (an invalid set)
	if answer <= 0 {
		return 1
	}

	// process the entire input array
	for i := 1; i < length; i++ {
		// check that the current index is not the same value as the previous item and that it increments by one
		// the first value is the final answer so end quickly here
		if A[i] != A[i-1]+1 && A[i] != A[i-1] {
			answer = A[i] - 1 // new answer
			return answer
		}
	}

	// backwards
	// because the challenge is looking for the smallest value it's more produent to start from the bottom and work your way to the top
	// if the challenge was asking for the largest positive integer then we'd want to go backwards
	/*
		for i := length - 2; i >= 0; i-- {
			fmt.Println(A[i])
			if A[i] != A[i+1]-1 && A[i] != A[i+1] {
				answer = A[i] + 1 // new answer
			}
		}
	*/

	// if, after processing the entire array, there isn't a change then the initialization case is the answer
	return answer
}
