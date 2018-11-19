package rand4go

func RandInt() int {
	return newRand().Int()
}

func RandIntn(n int) int {
	return newRand().Intn(n)
}
