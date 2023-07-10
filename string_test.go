package nrand_test

import (
	"fmt"
	"github.com/smartwalle/nrand"
	"testing"
)

func TestRandString(t *testing.T) {
	fmt.Println(nrand.String(6))
	fmt.Println(nrand.String(6))
	fmt.Println(nrand.String(6))
	fmt.Println(nrand.String(6))
	fmt.Println(nrand.String(6))
	fmt.Println(nrand.String(6))
	fmt.Println(nrand.String(6))
}

func BenchmarkString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		nrand.String(6)
	}
}

func BenchmarkRandString_Next(b *testing.B) {
	var rs = nrand.NewRandString(nrand.RandSourceAll)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rs.Next(6)
	}
}
