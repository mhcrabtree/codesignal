package challenges

import "sort"

// MakeArrayConsecutive Find the number of elements in array that are needed to make the array consecutive
func MakeArrayConsecutive(statues []int) int {
	sort.Ints(statues)
	needed := 0
	for i := 1; i < len(statues); i++ {
		temp := statues[i-1] + 1
		if statues[i] > temp {
			needed += (statues[i] - temp)
		}
	}
	return needed
}
