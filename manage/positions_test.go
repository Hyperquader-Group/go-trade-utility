package manage_test

import (
	"fmt"
	"testing"

	"github.com/go-numb/go-trade-utility/manage"
)

func TestCheckOrderSize(t *testing.T) {
	min := 0.06
	max := 3.0
	tension := 0.1

	s := manage.NewPosition(min, max)
	var has float64
	var count = 50
	for i := 0; i < count; i++ {
		has -= min
		s.Set(has)

		// full, size := s.Lot(1, tension)
		// fmt.Printf("buy:	%t,	%f\n", full, size)
		full, size := s.Lot(1, tension)
		fmt.Printf("sell:	%f,	%t,	%f\n", has, full, size)
	}
}
