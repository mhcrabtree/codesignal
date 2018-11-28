package challenges

// ShapeArea Find the area of a polygon for a given n.
func ShapeArea(n int) int {
	if n == 1 {
		return 1
	}
	return (n * n) + ((n - 1) * (n - 1))
}
