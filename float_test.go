package nrand_test

import (
	"fmt"
	"github.com/smartwalle/nrand"
	"testing"
)

func TestFloat32(t *testing.T) {
	fmt.Println(nrand.Float32())
	fmt.Println(nrand.Float32())
	fmt.Println(nrand.Float32())
	fmt.Println(nrand.Float32())
	fmt.Println(nrand.Float32())
}

func TestFloat64(t *testing.T) {
	fmt.Println(nrand.Float64())
	fmt.Println(nrand.Float64())
	fmt.Println(nrand.Float64())
	fmt.Println(nrand.Float64())
	fmt.Println(nrand.Float64())
}
