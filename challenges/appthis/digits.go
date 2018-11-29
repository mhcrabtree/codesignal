package challenges

import (
	"math"
	"sort"
)

// Digits calculates the highest number from an input of n digits where two numbers are created from the n digits w/o repeating a digit
func Digits(A []int) int {
	length := len(A)
	if length%2 == 0 {
		return digitsEven(A)
	}
	return digitsOdd(A)
}

// digitsEven the digits function to be run when the number of elements is even
// - this is faster than the digitsOdd function
func digitsEven(A []int) int {
	length := len(A)

	// sort ascending
	sort.Ints(A)

	// slice out evens
	var num0 []int

	// slice out odds
	var num1 []int

	// build numbers via iteration using power
	for i := 0; i < length/2; i++ {
		pow := int(math.Pow10(i))         // these will always be integers in this case (cast)
		num0 = append(num0, A[i*2]*pow)   // "common core" the even numbers
		num1 = append(num1, A[i*2+1]*pow) // "common core" the odd numbers
	}

	// sum the two numbers
	return sum([]int{sum(num0), sum(num1)})
}

// digitsOdd the digits function to be run when the number of elements is odd
// - this is slower than digitsEven
func digitsOdd(A []int) int {
	length := len(A)

	// use the original method if there's an even number of integers (it's faster)
	if length%2 == 0 {
		return Digits(A)
	}

	half := int(math.Ceil(float64(length) / 2.0))

	// sort ascending
	sort.Ints(A)

	// slice out evens
	var num0 []int

	// slice out odds
	var num1 []int

	// build numbers via iteration using power
	// iterate in reverse
	for i := half - 1; i >= 0; i-- {
		num0 = append(num0, A[i*2])
		if i*2-1 >= 0 {
			num1 = append(num1, A[i*2-1])
		}
	}

	a := arrayToNumber(num0)
	b := arrayToNumber(num1)

	// sum the two numbers
	return sum([]int{a, b})
}

// arrayToNumber will take an array of digits and create a numerical representation using math
func arrayToNumber(A []int) (s int) {
	length := len(A)
	for i, v := range A {
		pow := int(math.Pow10(length - 1 - i))
		s += v * pow
	}
	return s
}

// sums the integers inside of an array
func sum(a []int) (s int) {
	for _, v := range a {
		s += v
	}
	return
}
