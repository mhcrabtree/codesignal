package challenges

import (
	"math"
	"sort"
)

type Bucket struct {
	Digits []int
}

// ToNum sums the integers inside of an array
func (b *Bucket) ToNum() (s int) {
	for _, v := range b.Digits {
		s += v
	}
	return
}

// DigitsMod2 calculates the highest number from an input of n digits where two numbers are created from the n digits w/o repeating a digit
// A[] array of digits, unsorted
// n number of "numbers" to create from the A[] of digits
func DigitsMod2(A []int, n int) int {
	sort.Ints(A)

	nums := make([]Bucket, n)

	for i := 0; i < len(A); i++ {
		r := i % n                                       // get the remainder to determine into which bucket this value will go
		e := int(math.Ceil(float64(i+1)/float64(n))) - 1 // determine which exponent to use for Pow10
		nums[r].Digits = append(nums[r].Digits, A[i]*int(math.Pow10(e)))
	}

	// sum the values together
	s := 0
	for _, v := range nums {
		s += v.ToNum()
	}

	//fmt.Println(nums, s)
	return s
}
