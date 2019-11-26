package price

import (
	"fmt"
	"sync"
	"time"

	"github.com/mxmCherry/movavg"
)

type Source struct {
	Long  Prices
	Short Prices
}

type Prices struct {
	Term    int
	Prices  []float64
	Volumes []float64
}

func SMA(s *Source) float64 {
	start := time.Now()
	defer func() {
		end := time.Now()
		fmt.Println("exec time: ", end.Sub(start))
	}()
	maLong := movavg.NewSMA(s.Long.Term)
	maShort := movavg.NewSMA(s.Short.Term)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for _, f := range s.Long.Prices {
			maLong.Add(f)
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for _, f := range s.Short.Prices {
			maShort.Add(f)
		}
		wg.Done()
	}()

	wg.Wait()

	return maShort.Avg() / maLong.Avg()
}
