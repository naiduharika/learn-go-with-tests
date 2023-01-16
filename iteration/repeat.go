package iteration

const repeatCount = 5

// Repeat returns character repeated repeatCount times.
func Repeat(character string, count int) string {

	if count == 0 {
		count = repeatCount
	}

	var repeated string
	for i := 0; i < count; i++ {
		repeated += character
	}
	return repeated
}
