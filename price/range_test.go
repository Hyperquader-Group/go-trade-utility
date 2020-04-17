package price_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/go-numb/go-trade-utility/price"
)

func TestSetToRange(t *testing.T) {
	// 2桁100単位でkeyを作り、同じ価格帯注文価格をチェックする
	prange := price.NewRange(1, 2)

	count := 1000
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for i := 0; i < count; i++ {
		order := 6000.0 + float64(r.Intn(1000))
		var t bool

		if i%2 == 0 {
			t = prange.Set(1, order)
			fmt.Printf("%t %.1f\n", t, order)
		} else {
			t = prange.Set(1, order)
			fmt.Printf("%t %.1f\n", t, order)
		}

		if t {
			// 注文履歴リセット
			prange.Reset()
		}
	}
}
