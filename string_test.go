package rand4go

import (
	"fmt"
	"testing"
)

func TestRandString(t *testing.T) {
	fmt.Println(String(6, RandTypeDefault))
	fmt.Println(String(6, RandTypeNumOnly))
	fmt.Println(String(6, RandTypeLowerOnly))
	fmt.Println(String(6, RandTypeUpperOnly))
	fmt.Println(String(6, RandTypeLowerNum))
	fmt.Println(String(6, RandTypeUpperNum))
	fmt.Println(String(6, RandTypeLowerUpper))
}
