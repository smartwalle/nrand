package rand4go

func Int() int {
	return newRand().Int()
}

func Intn(n int) int {
	return newRand().Intn(n)
}

func IntRange(min, max int) int {
	if min == max {
		return min
	}
	if min > max {
		min, max = max, min
	}
	return newRand().Intn(max-min) + min
}
