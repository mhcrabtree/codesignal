package models

// DataIntInt struct for holding int array and output of single int value
type DataIntInt struct {
	Input  int  // the data used for the test
	Output int  // the correct answer for the test data
	Test   *int // the value being tested
	Pass   bool // the return status of pass or fail
}

// Assert returns true or false depending on if the value passed in matches the object output value
func (d *DataIntInt) Assert(v *int) bool {
	d.Test = v // save the most recent value to test against the correct output answer
	if *v == d.Output {
		d.Pass = true
	}
	return d.Pass
}
