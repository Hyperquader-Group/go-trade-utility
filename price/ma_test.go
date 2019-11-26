package price

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"sort"
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
