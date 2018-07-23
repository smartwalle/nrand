package str4go

import (
	"math/rand"
	"time"
)

const (
	k_SOURCE = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
)

func RandString(size int) (str string) {
	var r = rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < size; i++ {
		idx := r.Intn(len(k_SOURCE) - 1)
		str += k_SOURCE[idx : idx+1]
	}
	return str
}
