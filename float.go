package nrand

func Float32() float32 {
	return sharedRand().Float32()
}

func Float64() float64 {
	return sharedRand().Float64()
}
