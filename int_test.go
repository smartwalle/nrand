package rand4go

import (
	"fmt"
	"testing"
)

func TestRandInt(t *testing.T) {
	fmt.Println(Int())
	fmt.Println(Int())
	fmt.Println(Int())
	fmt.Println(Int())
	fmt.Println(Int())
}

func TestRandIntn(t *testing.T) {
	fmt.Println(Intn(10))
	fmt.Println(Intn(20))
	fmt.Println(Intn(30))
	fmt.Println(Intn(40))
	fmt.Println(Intn(50))
}

func TestRandIntRange(t *testing.T) {
	fmt.Println(IntRange(3, 10))
	fmt.Println(IntRange(3, 10))
	fmt.Println(IntRange(3, 10))
	fmt.Println(IntRange(3, 10))
	fmt.Println(IntRange(3, 10))
	fmt.Println(IntRange(1, 3))
}
