package pointer

import (
	"fmt"
	"testing"
)

func Test_Pointer(t *testing.T) {
	var a *int
	a = new(int)
	*a = 100
	fmt.Println(*a)
}
