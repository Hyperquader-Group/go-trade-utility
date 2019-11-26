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
	f, _ := os.OpenFile("./f.csv", os.O_CREATE|os.O_WRONLY, 0755)
	defer f.Close()

	w := csv.NewWriter(f)

	count := 200
	ff := make([]float64, count)
	for i := 0; i < count; i++ {
		rand.Seed(time.Now().UnixNano())
		f := rand.Float64() * 10
		ff[i] = f

		time.Sleep(time.Microsecond)
	}

	sort.Float64s(ff)

	for i := 0; i < count; i++ {
		w.Write([]string{fmt.Sprintf("%f", ff[i])})
	}
	w.Flush()

	s := &Source{}
	s.Long.Term = 21
	s.Long.Prices = ff
	s.Short.Term = 9
	s.Short.Prices = ff

	diverd := SMA(s)
	fmt.Printf("%f\n", diverd)
}
