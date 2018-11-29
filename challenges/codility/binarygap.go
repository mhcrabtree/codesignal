package challenges

// BinaryGap Find longest sequence of zeros in binary representation of an integer.
func BinaryGap(n int) int {
	//var b []int
	hasStarted := false
	workingZeroCount := 0
	largestZeroCount := 0
	for n > 0 {
		t := n % 2
		//b = append(b, t)
		//fmt.Println(n, t)

		// init counter
		if !hasStarted && t == 1 {
			hasStarted = true
		}

		if hasStarted && t == 0 {
			workingZeroCount++
		} else if hasStarted { // possible reset
			if workingZeroCount > largestZeroCount {
				largestZeroCount = workingZeroCount
			}
			workingZeroCount = 0
		}

		n = n >> 1
	}

	//	fmt.Println(b, largestZeroCount)

	return largestZeroCount
}
