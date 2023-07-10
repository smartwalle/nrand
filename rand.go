package nrand

import (
	"math/rand"
	"time"
)

var shared = NewRand(time.Now().UnixNano())

type Rand struct {
	r *rand.Rand
}

func NewRand(seed int64) *Rand {
	var r = &Rand{}
	r.r = rand.New(rand.NewSource(seed))
	return r
}

func (this *Rand) Int() int {
	return this.r.Int()
}

func (this *Rand) Intn(n int) int {
	return this.r.Intn(n)
}

func (this *Rand) IntRange(min, max int) int {
	if min == max {
		return min
	}
	if min > max {
		min, max = max, min
	}
	return this.Intn(max-min) + min
}

func (this *Rand) Int32() int32 {
	return this.r.Int31()
}

func (this *Rand) Int32n(n int32) int32 {
	return this.r.Int31n(n)
}

func (this *Rand) Int32Range(min, max int32) int32 {
	if min == max {
		return min
	}
	if min > max {
		min, max = max, min
	}
	return this.Int32n(max-min) + min
}

func (this *Rand) Int64() int64 {
	return this.r.Int63()
}

func (this *Rand) Int64n(n int64) int64 {
	return this.r.Int63n(n)
}

func (this *Rand) Int64Range(min, max int64) int64 {
	if min == max {
		return min
	}
	if min > max {
		min, max = max, min
	}
	return this.Int64n(max-min) + min
}

func (this *Rand) Uint32() uint32 {
	return this.r.Uint32()
}

func (this *Rand) Uint64() uint64 {
	return this.r.Uint64()
}

func (this *Rand) Float32() float32 {
	return this.r.Float32()
}

func (this *Rand) Float64() float64 {
	return this.r.Float64()
}
