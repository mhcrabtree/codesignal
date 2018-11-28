package challenges

import (
	"fmt"
	"sort"
)

type Bucket struct {
	Digits []int
}

// DigitsMod2 calculates the highest number from an input of n digits where two numbers are created from the n digits w/o repeating a digit
func DigitsMod2(A []int, n int) int {
	//	length := len(A)
	sort.Ints(A)

	//	var a [3]Bucket
	nums := make([]Bucket, n)

	fmt.Println(nums)
	return 0
}
