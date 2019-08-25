package repeat

const totalRepeat = 5

// Repeat character 5 times
func Repeat(character string) string {
	var result string
	for i := 0; i < totalRepeat; i++ {
		result += character
	}

	return result
}
