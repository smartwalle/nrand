package rand4go

// --------------------------------------------------------------------------------
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
	return Intn(max-min) + min
}

// --------------------------------------------------------------------------------
func Int32() int32 {
	return newRand().Int31()
}

func Int32n(n int32) int32 {
	return newRand().Int31n(n)
}

func Int32Range(min, max int32) int32 {
	if min == max {
		return min
	}
	if min > max {
		min, max = max, min
	}
	return Int32n(max-min) + min
}

// --------------------------------------------------------------------------------
func Int64() int64 {
	return newRand().Int63()
}

func Int64n(n int64) int64 {
	return newRand().Int63n(n)
}

func Int64Range(min, max int64) int64 {
	if min == max {
		return min
	}
	if min > max {
		min, max = max, min
	}
	return Int64n(max-min) + min
}

// --------------------------------------------------------------------------------
func UInt32() uint32 {
	return newRand().Uint32()
}

func UInt64() uint64 {
	return newRand().Uint64()
}
