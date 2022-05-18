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

func String(size int, source RandSource) string {
	var rs = NewRandString(source)
	return rs.Next(size)
}

type RandString struct {
	rs  rand.Source
	src []byte
}

func NewRandString(source RandSource) *RandString {
	var rs = &RandString{}
	rs.rs = rand.NewSource(time.Now().UnixNano())

	var upperChar = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var lowerChar = "abcdefghijklmnopqrstuvwxyz"
	var number = "0123456789"

	switch source {
	case RandSourceAll:
		rs.src = []byte(upperChar + lowerChar + number)
	case RandSourceNum:
		rs.src = []byte(number)
	case RandSourceLower:
		rs.src = []byte(lowerChar)
	case RandSourceUpper:
		rs.src = []byte(upperChar)
	case RandSourceLowerNum:
		rs.src = []byte(lowerChar + number)
	case RandSourceUpperNum:
		rs.src = []byte(upperChar + number)
	case RandSourceLowerUpper:
		rs.src = []byte(upperChar + lowerChar)
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
