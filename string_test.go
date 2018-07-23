package str4go

import (
	"fmt"
	"testing"
)

func TestRandString(t *testing.T) {
	fmt.Println(RandString(6, K_RAND_TYPE_DEFAULT))
	fmt.Println(RandString(6, K_RAND_TYPE_NUM_ONLY))
	fmt.Println(RandString(6, K_RAND_TYPE_LOWER_ONLY))
	fmt.Println(RandString(6, K_RAND_TYPE_UPPER_ONLY))
	fmt.Println(RandString(6, K_RAND_TYPE_LOWER_NUM))
	fmt.Println(RandString(6, K_RAND_TYPE_UPPER_NUM))
	fmt.Println(RandString(6, K_RAND_TYPE_LOWER_UPPER))
}
