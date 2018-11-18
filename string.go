package rand4go

const (
	k_SOURCE_UPPER = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	k_SOURCE_LOWER = "abcdefghijklmnopqrstuvwxyz"
	k_SOURCE_NUM   = "0123456789"
)

const (
	K_RAND_TYPE_DEFAULT     = 0
	K_RAND_TYPE_NUM_ONLY    = 1
	K_RAND_TYPE_LOWER_ONLY  = 2
	K_RAND_TYPE_UPPER_ONLY  = 3
	K_RAND_TYPE_LOWER_NUM   = 4
	K_RAND_TYPE_UPPER_NUM   = 5
	K_RAND_TYPE_LOWER_UPPER = 6
)

var source = make(map[int]string)

func init() {
	source[K_RAND_TYPE_DEFAULT] = k_SOURCE_UPPER + k_SOURCE_LOWER + k_SOURCE_NUM
	source[K_RAND_TYPE_NUM_ONLY] = k_SOURCE_NUM
	source[K_RAND_TYPE_LOWER_ONLY] = k_SOURCE_LOWER
	source[K_RAND_TYPE_UPPER_ONLY] = k_SOURCE_UPPER
	source[K_RAND_TYPE_LOWER_NUM] = k_SOURCE_LOWER + k_SOURCE_NUM
	source[K_RAND_TYPE_UPPER_NUM] = k_SOURCE_UPPER + k_SOURCE_NUM
	source[K_RAND_TYPE_LOWER_UPPER] = k_SOURCE_UPPER + k_SOURCE_LOWER
}

func RandString(size, rType int) (str string) {
	var src = source[rType]
	if src == "" {
		src = source[K_RAND_TYPE_DEFAULT]
	}

	var r = newRand()
	for i := 0; i < size; i++ {
		idx := r.Intn(len(src) - 1)
		str += src[idx : idx+1]
	}
	return str
}
