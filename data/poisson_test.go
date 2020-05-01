package data_test

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/go-numb/go-trade-utility/data"
)

func TestPoisson(t *testing.T) {
	f, _ := os.Open("/Volumes/DailySD/SD_Desktop/diff.csv")
	defer f.Close()

	r := csv.NewReader(f)

	// 期間内板置きで約定率10%を欲するときの約定見込み指値幅を算出
	s := data.NewValues()
	// for i := 0; i < count; i++ {
	// 	s.Set(float64(rand.Intn(1000)))
	// }

	for {
		row, err := r.Read()
		if err != nil {
			break
		}

		x, _ := strconv.ParseFloat(row[0], 64)
		s.Set(x)
	}

	start := time.Now()
	fmt.Printf("%+v\n", s.Threshold(true, 0.5))
	fmt.Printf("%+v\n", time.Since(start))
}
