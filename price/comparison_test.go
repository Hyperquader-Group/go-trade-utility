package price

import (
	"fmt"
	"testing"
)

func TestComparison(t *testing.T) {
	comp := NewComparison()

	count := 100
	// var add int
	for i := 0; i < count; i++ {
		// if i%2 == 0 {
		// 	add = -1
		// }
		comp.Set(float64((i * 10) + 6000))
		fmt.Printf("%+v\n", comp.Ratio())
	}
}
