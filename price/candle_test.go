package price_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-numb/go-ftx/rest"
	"github.com/go-numb/go-ftx/rest/public/markets"
	"github.com/go-numb/go-trade-utility/price"
	"github.com/stretchr/testify/assert"
)

func TestHistricals(t *testing.T) {
	c := rest.New(nil)

	candle := price.NewCandle()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	var count int
	start := time.Now()
	for {
		select {
		case <-ticker.C:
			trades, err := c.Trades(&markets.RequestForTrades{
				ProductCode: "BTC-PERP",
				Limit:       2,
			})
			assert.NoError(t, err)

			for _, v := range *trades {
				candle.Set(v.Price, v.Size)
				if 100 < candle.Volume() { // 反転足の発見
					count++
					ohlcv := candle.Hige(0.002) // 変動値幅閾値
					if ohlcv.IsSufficient {
						fmt.Printf("%d	%.1f	%s	%vBTC	%v	%v\n", count, v.Price, ohlcv, candle.Volume(), time.Now().Sub(start), time.Now())
					}
					candle.Reset()
				}
			}
		}
	}
}
