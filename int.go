package rand4go

func Int() int {
	return sharedRand().Int()
}

func Intn(n int) int {
	return sharedRand().Intn(n)
}

func IntRange(min, max int) int {
	return sharedRand().IntRange(min, max)
}

func Int32() int32 {
	return sharedRand().Int32()
}

func Int32n(n int32) int32 {
	return sharedRand().Int32n(n)
}

func Int32Range(min, max int32) int32 {
	return sharedRand().Int32Range(min, max)
}

func Int64() int64 {
	return sharedRand().Int64()
}

func Int64n(n int64) int64 {
	return sharedRand().Int64n(n)
}

func Int64Range(min, max int64) int64 {
	return sharedRand().Int64Range(min, max)
}

func Uint32() uint32 {
	return sharedRand().Uint32()
}

func Uint64() uint64 {
	return sharedRand().Uint64()
}
