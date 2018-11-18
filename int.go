package rand4go

import "math/rand"

func RandInt() int {
	return newRand().Int()
}

func RandIntn(n int) int {
	return newRand().Intn(n)
	return rand.Int()
}
