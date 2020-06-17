package rand4go

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

var sources = make(map[RandSource][]rune)

func init() {
	sources[RandSourceAll] = []rune(kUpperChar + kLowerChar + kNumber)
	sources[RandSourceNum] = []rune(kNumber)
	sources[RandSourceLower] = []rune(kLowerChar)
	sources[RandSourceUpper] = []rune(kUpperChar)
	sources[RandSourceLowerNum] = []rune(kLowerChar + kNumber)
	sources[RandSourceUpperNum] = []rune(kUpperChar + kNumber)
	sources[RandSourceLowerUpper] = []rune(kUpperChar + kLowerChar)
}

func String(size int, source RandSource) string {
	var s = sources[source]
	if len(s) == 0 {
		s = sources[RandSourceAll]
	}

	var d = make([]rune, size)
	var r = newRand()

	for i := range d {
		d[i] = s[r.Intn(len(s))]
	}
	return string(d)
}
