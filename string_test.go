package rand4go

import (
	"fmt"
	"testing"
)

func TestRandString(t *testing.T) {
	fmt.Println(String(6, K_RAND_TYPE_DEFAULT))
	fmt.Println(String(6, K_RAND_TYPE_NUM_ONLY))
	fmt.Println(String(6, K_RAND_TYPE_LOWER_ONLY))
	fmt.Println(String(6, K_RAND_TYPE_UPPER_ONLY))
	fmt.Println(String(6, K_RAND_TYPE_LOWER_NUM))
	fmt.Println(String(6, K_RAND_TYPE_UPPER_NUM))
	fmt.Println(String(6, K_RAND_TYPE_LOWER_UPPER))
}
