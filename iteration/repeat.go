package iteration

func Repeat(a string) string {
	var repeat string
	for i := 0; i < 7; i++ {
		repeat += a
	}
	return repeat
}
