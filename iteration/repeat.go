package iteration

const repeatCount = 5

// Returns character repeated 5 times
func RepeatFiveTimes(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}
