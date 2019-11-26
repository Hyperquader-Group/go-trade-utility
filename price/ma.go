package price

import (
	"errors"
	"sync"

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

func SMA(isVWAP bool, s *Source) (div float64, err error) {
	if isVWAP {
		if s.Long.Volumes == nil {
			return 0, errors.New("vwap volume is nil")
		} else if s.Short.Volumes == nil {
			return 0, errors.New("vwap volume is nil")
		}
		if len(s.Long.Prices) != len(s.Long.Volumes) {
			return 0, errors.New("vwap long price & volume length dont match")
		} else if len(s.Short.Prices) != len(s.Short.Volumes) {
			return 0, errors.New("vwap short price & volume length dont match")
		}
	}
	maLong := movavg.NewSMA(s.Long.Term)
	maShort := movavg.NewSMA(s.Short.Term)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		if !isVWAP {
			for _, f := range s.Long.Prices {
				maLong.Add(f)
			}
		} else { // 出来高加重平均
			for i, f := range s.Long.Prices {
				maLong.Add(f * s.Long.Volumes[i])
			}
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		if !isVWAP {
			for _, f := range s.Short.Prices {
				maShort.Add(f)
			}
		} else { // 出来高加重平均
			for i, f := range s.Short.Prices {
				maShort.Add(f * s.Short.Volumes[i])
			}
		}
		wg.Done()
	}()

	wg.Wait()

	return maShort.Avg() / maLong.Avg(), nil
}
