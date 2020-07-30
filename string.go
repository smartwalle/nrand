package rand4go

import "strings"

const (
	kUpperChar = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	kLowerChar = "abcdefghijklmnopqrstuvwxyz"
	kNumber    = "0123456789"
)

type RandSource int

const (
	RandSourceAll        RandSource = 0
	RandSourceNum        RandSource = 1
	RandSourceLower      RandSource = 2
	RandSourceUpper      RandSource = 3
	RandSourceLowerNum   RandSource = 4
	RandSourceUpperNum   RandSource = 5
	RandSourceLowerUpper RandSource = 6
)

var sources = make(map[RandSource][]byte)

func init() {
	sources[RandSourceAll] = []byte(kUpperChar + kLowerChar + kNumber)
	sources[RandSourceNum] = []byte(kNumber)
	sources[RandSourceLower] = []byte(kLowerChar)
	sources[RandSourceUpper] = []byte(kUpperChar)
	sources[RandSourceLowerNum] = []byte(kLowerChar + kNumber)
	sources[RandSourceUpperNum] = []byte(kUpperChar + kNumber)
	sources[RandSourceLowerUpper] = []byte(kUpperChar + kLowerChar)
}

func String(size int, source RandSource) string {
	var s = sources[source]
	if len(s) == 0 {
		s = sources[RandSourceAll]
	}

	var r = newRand()

	var b = strings.Builder{}
	b.Grow(size)

	for i := 0; i < size; i++ {
		b.WriteByte(s[r.Intn(len(s))])
	}
	return b.String()
}
