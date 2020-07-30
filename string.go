package rand4go

import (
	"math/rand"
	"strings"
	"time"
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
	r   *rand.Rand
	src []byte
}

func NewRandString(source RandSource) *RandString {
	var rs = &RandString{}
	rs.r = rand.New(rand.NewSource(time.Now().UnixNano()))

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

func (this *RandString) Next(size int) string {
	var b = strings.Builder{}
	b.Grow(size)

	for i := 0; i < size; i++ {
		b.WriteByte(this.src[this.r.Intn(len(this.src))])
	}
	return b.String()
}
