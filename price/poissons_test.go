package price_test

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	"gonum.org/v1/gonum/stat/distuv"

	"github.com/go-numb/go-trade-utility/price"
)

func TestPoissons(t *testing.T) {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	// 単位時間1回出現
	term := 3 * time.Second
	p := &distuv.Poisson{
		Lambda: 1,
	}
	po := &price.Poissons{
		Term:    term,
		Appears: make([]float64, 0),
		Po:      p,
	}

	ticker := time.NewTicker(term)
	defer ticker.Stop()

	count := 20
	min := 50.0

	for {
		select {
		case <-ticker.C:
			for i := 0; i < count; i++ {
				appeared := (3000.0 * p.Prob(float64(r.Intn(11)))) + min
				po.Appears = append(po.Appears, appeared)
				// fmt.Printf("%+v	-> %.2f\n", appeared, po.Po.Prob(appeared)*100)
				if err := save(po.Appears); err != nil {
					t.Fatal(err)
				}
				po.Appears = []float64{}
			}

		}
	}

}

func save(fx []float64) error {
	f, err := os.OpenFile("./spread.csv", os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)

	for i := range fx {
		w.Write([]string{fmt.Sprintf("%v", fx[i])})
	}
	w.Flush()

	return nil
}
