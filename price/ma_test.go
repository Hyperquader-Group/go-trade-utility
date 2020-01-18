package price

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"testing"
	"time"
)

func TestMa(t *testing.T) {
	f, _ := os.OpenFile("./f.csv", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	defer f.Close()

	w := csv.NewWriter(f)

	count := 200
	ff := make([]float64, count)
	vv := make([]float64, count)
	for i := 0; i < count; i++ {
		rand.Seed(time.Now().UnixNano())
		f := rand.Float64() * 10
		ff[i] = f
		vv[i] = f * 0.01
		time.Sleep(time.Microsecond)
	}

	sort.Float64s(ff) // 降順
	// sort.Sort(sort.Reverse(sort.Float64Slice(ff))) // 昇順

	for i := 0; i < count; i++ {
		w.Write([]string{fmt.Sprintf("%f", ff[i])})
	}
	w.Flush()

	s := &Source{}
	s.Long.Term = 21
	s.Long.Prices = ff
	s.Long.Volumes = vv
	s.Short.Term = 9
	s.Short.Prices = ff
	s.Short.Volumes = vv

	div, err := SMA(true, s)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%f\n", div)
}

// 異常価格の分布を取得
// 1（乖離なし）付近はDrop()し、正規分布に整形、分足レベルでのRektInを検出する
func TestAnalytics(t *testing.T) {
	f, err := os.Open("/Users/<name>/project_bitcoin/OHLC_chart/bitstampUSD_1-min_data_2012-01-01_to_2018-03-27.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	results, err := r.ReadAll()
	if err != nil {
		t.Fatal(err)
	}

	prices := make([]float64, len(results))
	volumes := make([]float64, len(results))
	for i, row := range results {
		price, err := strconv.ParseFloat(row[1], 64)
		if err != nil {
			continue
		}
		volume, err := strconv.ParseFloat(row[5], 64)
		if err != nil {
			continue
		}
		prices[i] = price
		volumes[i] = volume
	}

	fmt.Printf("%+v\n", prices[:10])
	fmt.Printf("%+v\n", volumes[:10])

	saveFile, err := os.OpenFile("./divers.csv", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		t.Fatal(err)
	}
	defer saveFile.Close()

	wCSV := csv.NewWriter(saveFile)

	for i := range prices {
		s := &Source{}
		s.Long.Term = 21
		s.Long.Prices = prices[i : 21+i]
		s.Long.Volumes = volumes[i : 21+i]
		s.Short.Term = 9
		s.Short.Prices = prices[i : 21+i]
		s.Short.Volumes = volumes[i : 21+i]

		div, err := SMA(true, s)
		if err != nil {
			t.Fatal(err)
		}

		wCSV.Write([]string{fmt.Sprintf("%f", div)})

		if len(prices)-22 < i {
			break
		}
	}
	wCSV.Flush()
}
