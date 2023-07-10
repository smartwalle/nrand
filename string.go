package nrand

import (
	"math/rand"
	"time"
	"unsafe"
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

const (
	kUpperChar = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	kLowerChar = "abcdefghijklmnopqrstuvwxyz"
	kNumber    = "0123456789"
)

var rString = NewRandString(RandSourceAll)

func String(size int) string {
	return rString.Next(size)
}

type RandString struct {
	rs  rand.Source
	src []byte
}

func NewRandString(source RandSource) *RandString {
	var rs = &RandString{}
	rs.rs = rand.NewSource(time.Now().UnixNano())

	switch source {
	case RandSourceAll:
		rs.src = []byte(kUpperChar + kLowerChar + kNumber)
	case RandSourceNum:
		rs.src = []byte(kNumber)
	case RandSourceLower:
		rs.src = []byte(kLowerChar)
	case RandSourceUpper:
		rs.src = []byte(kUpperChar)
	case RandSourceLowerNum:
		rs.src = []byte(kLowerChar + kNumber)
	case RandSourceUpperNum:
		rs.src = []byte(kUpperChar + kNumber)
	case RandSourceLowerUpper:
		rs.src = []byte(kUpperChar + kLowerChar)
	}
	return rs
}

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// Next https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
func (this *RandString) Next(size int) string {
	var b = make([]byte, size)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := size-1, this.rs.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = this.rs.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(this.src) {
			b[i] = this.src[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return *(*string)(unsafe.Pointer(&b))
}
