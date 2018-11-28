package challenges

// ZigZag For a given array of integers return the length of its longest contiguous sub-array that is a zigzag sequence
func ZigZag(a []int) int {
	length := len(a)

	if len(a) == 1 {
		return length
	}

	dir := ""
	var workingSeq []int                  // this will hold the working set
	var longestSeq []int                  // this will hold the first longest set we've found
	workingSeq = append(workingSeq, a[0]) // prime the sequence with the first value

	for i := 1; i < length; i++ {
		seqLen := len(workingSeq)

		if a[i] > workingSeq[seqLen-1] && dir != "down" {
			dir = "down"
		} else if a[i] < workingSeq[seqLen-1] && dir != "up" {
			dir = "up"
		} else {
			if len(longestSeq) < len(workingSeq) {
				//fmt.Println(longestSeq, "vs", workingSeq)
				longestSeq = workingSeq
			}
			workingSeq = nil // empty the slice
			if a[i-1] != a[i] {
				workingSeq = append(workingSeq, a[i-1]) // append the previous item because it's still valid so long as it's not the same as the current item
			}
		}

		workingSeq = append(workingSeq, a[i])
	}
	//fmt.Println("Answer", longestSeq)

	return len(longestSeq)
}
