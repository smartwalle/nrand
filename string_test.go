package rand4go

import (
	"fmt"
	"testing"
)

func TestRandString(t *testing.T) {
	fmt.Println(String(6, RandSourceAll))
	fmt.Println(String(6, RandSourceNum))
	fmt.Println(String(6, RandSourceLower))
	fmt.Println(String(6, RandSourceUpper))
	fmt.Println(String(6, RandSourceLowerNum))
	fmt.Println(String(6, RandSourceUpperNum))
	fmt.Println(String(6, RandSourceLowerUpper))
}

func BenchmarkString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		String(6, RandSourceAll)
	}
}
