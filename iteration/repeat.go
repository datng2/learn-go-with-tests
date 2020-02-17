package iteration

// Repeat takes a character and repeats it N times
func Repeat(character string, numRepeats int) string {
	var repeated string = ""
	for i := 0; i < numRepeats; i++ {
		repeated += character
	}
	return repeated
}
