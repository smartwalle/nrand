package rand4go

import (
	"math/rand"
	"time"
)

func newRand() *rand.Rand {
	var r = rand.New(rand.NewSource(time.Now().UnixNano()))
	return r
}
