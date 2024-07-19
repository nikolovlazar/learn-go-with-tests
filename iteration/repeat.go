package iteration

const defaultRepeatCount = 5

func Repeat(character string, repeatCount int) string {
	var repeated string

	repeat := defaultRepeatCount
	if repeatCount > 0 {
		repeat = repeatCount
	}

	for i := 0; i < repeat; i++ {
		repeated += character
	}

	return repeated
}
