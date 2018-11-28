package challenges

// AdjacentElementsProduct Given an array of integers, find the pair of adjacent elements that has the largest product and return that product.
func AdjacentElementsProduct(inputArray []int) int {
	var product int
	for i := 1; i < len(inputArray); i++ {
		temp := inputArray[i-1] * inputArray[i]
		if product == 0 || temp > product {
			product = temp
		}
	}
	return product
}
