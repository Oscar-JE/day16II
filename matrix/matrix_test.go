package matrix

import (
	"fmt"
	"testing"
)

func TestUtskrift(t *testing.T) {
	mat := Init(2, 4)
	mat.Set(1, 2, 5)
	fmt.Println(mat)
}
