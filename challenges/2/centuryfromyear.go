package challenges

// CenturyFromYear Given a year, return the century it is in.
func CenturyFromYear(year int) int {
	century := year / 100
	if year%100 != 0 {
		century++
	}
	return century

	// same answer using the math api and type casting
	//	return int(math.Ceil(float64(year) / 100.0))
}
