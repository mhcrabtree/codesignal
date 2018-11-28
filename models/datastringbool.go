package models

// DataStringBool struct for holding string and output of single bool value
type DataStringBool struct {
	Input  string // the data used for the test
	Output bool   // the correct answer for the test data
	Test   *bool  // the value being tested
	Pass   bool   // the return status of pass or fail
}

// Assert returns true or false depending on if the value passed in matches the object output value
func (d *DataStringBool) Assert(v *bool) bool {
	d.Test = v // save the most recent value to test against the correct output answer
	if *v == d.Output {
		d.Pass = true
	}
	return d.Pass
}
