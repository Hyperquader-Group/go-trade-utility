package data

import (
	"math"
	"sort"
	"sync"
)

// Values 値を保持し、値の中から指定出現率の閾値を取得する
// ある約定率を得たい、場合を想定
type Values struct {
	mux sync.RWMutex

	length int
	values []float64
}

// NewValues ある期間に設定確率出現する場合
// 次の期間に設定したスプレッドで約定する確率
func NewValues() *Values {
	return &Values{
		values: make([]float64, 0),
	}
}

func (p *Values) Set(value float64) {
	p.mux.Lock()
	defer p.mux.Unlock()

	p.length++
	p.values = append(p.values, value)
}

func (p *Values) Len() int {
	p.mux.RLock()
	defer p.mux.RUnlock()
	return p.length
}

func (p *Values) ValueLen() int {
	p.mux.RLock()
	defer p.mux.RUnlock()
	return len(p.values)
}

func (p *Values) Reset() {
	p.mux.Lock()
	defer p.mux.Unlock()

	// n-1分は保持
	// Reset直後の0値を除くため
	if len(p.values) == p.length {
		p.length = 0
		return
	}
	// 先入れ後出しで減らしていく
	p.values = p.values[p.length:]
	p.length = 0

}

func (p *Values) Threshold(isReverse bool, prob float64) float64 {
	p.mux.RLock()
	defer p.mux.RUnlock()
	if len(p.values) < 1 {
		return 0
	}

	fx := p.values
	if isReverse {
		sort.Sort(sort.Reverse(sort.Float64Slice(fx)))
	} else {
		sort.Sort(sort.Float64Slice(fx))
	}

	// mean, stdv := stat.MeanStdDev(p.values, nil)
	// ex := stat.ExKurtosis(p.values, nil)
	// fmt.Printf("%+v	%v	%v\n", mean, stdv, ex)

	// 累積確率から得たい出現回数の余事象を捨てる
	point := int(math.Floor(float64(len(fx)) * prob))
	if len(fx) <= point+1 {
		return 0
	}
	return fx[point+1]
}

func (p *Values) Copy() []float64 {
	p.mux.RLock()
	defer p.mux.RUnlock()

	return p.values
}
