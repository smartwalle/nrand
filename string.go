package rand4go

const (
	kUpperChar = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	kLowerChar = "abcdefghijklmnopqrstuvwxyz"
	kNumber    = "0123456789"
)

const (
	RandTypeDefault    = 0
	RandTypeNumOnly    = 1
	RandTypeLowerOnly  = 2
	RandTypeUpperOnly  = 3
	RandTypeLowerNum   = 4
	RandTypeUpperNum   = 5
	RandTypeLowerUpper = 6
)

var source = make(map[int]string)

func init() {
	source[RandTypeDefault] = kUpperChar + kLowerChar + kNumber
	source[RandTypeNumOnly] = kNumber
	source[RandTypeLowerOnly] = kLowerChar
	source[RandTypeUpperOnly] = kUpperChar
	source[RandTypeLowerNum] = kLowerChar + kNumber
	source[RandTypeUpperNum] = kUpperChar + kNumber
	source[RandTypeLowerUpper] = kUpperChar + kLowerChar
}

func String(size, rType int) (str string) {
	var src = source[rType]
	if src == "" {
		src = source[RandTypeDefault]
	}

	var r = newRand()
	for i := 0; i < size; i++ {
		idx := r.Intn(len(src) - 1)
		str += src[idx : idx+1]
	}
	return str
}
