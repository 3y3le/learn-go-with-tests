package iteration

func Repeat(buffer string, count int) (context string) {
	for i := 0; i < count; i++ {
		context += buffer
	}
	return
}
