package nrand_test

import (
	"fmt"
	"github.com/smartwalle/nrand"
	"testing"
)

func TestRandInt(t *testing.T) {
	fmt.Println(nrand.Int())
	fmt.Println(nrand.Int())
	fmt.Println(nrand.Int())
	fmt.Println(nrand.Int())
	fmt.Println(nrand.Int())
}

func TestRandIntn(t *testing.T) {
	fmt.Println(nrand.Intn(10))
	fmt.Println(nrand.Intn(20))
	fmt.Println(nrand.Intn(30))
	fmt.Println(nrand.Intn(40))
	fmt.Println(nrand.Intn(50))
}

func TestRandIntRange(t *testing.T) {
	fmt.Println(nrand.IntRange(3, 10))
	fmt.Println(nrand.IntRange(3, 10))
	fmt.Println(nrand.IntRange(3, 10))
	fmt.Println(nrand.IntRange(3, 10))
	fmt.Println(nrand.IntRange(3, 10))
	fmt.Println(nrand.IntRange(1, 3))
}
