package nrand

func Int() int {
	return shared.Int()
}

func Intn(n int) int {
	return shared.Intn(n)
}

func IntRange(min, max int) int {
	return shared.IntRange(min, max)
}

func Int32() int32 {
	return shared.Int32()
}

func Int32n(n int32) int32 {
	return shared.Int32n(n)
}

func Int32Range(min, max int32) int32 {
	return shared.Int32Range(min, max)
}

func Int64() int64 {
	return shared.Int64()
}

func Int64n(n int64) int64 {
	return shared.Int64n(n)
}

func Int64Range(min, max int64) int64 {
	return shared.Int64Range(min, max)
}

func Uint32() uint32 {
	return shared.Uint32()
}

func Uint64() uint64 {
	return shared.Uint64()
}
