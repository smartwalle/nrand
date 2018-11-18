package rand4go

import (
	"fmt"
	"testing"
)

func TestRandInt(t *testing.T) {
	fmt.Println(RandInt())
	fmt.Println(RandInt())
	fmt.Println(RandInt())
	fmt.Println(RandInt())
}

func TestRandIntn(t *testing.T) {
	fmt.Println(RandIntn(10))
	fmt.Println(RandIntn(20))
	fmt.Println(RandIntn(30))
	fmt.Println(RandIntn(40))
	fmt.Println(RandIntn(50))
}
