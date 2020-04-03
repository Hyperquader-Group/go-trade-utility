package price

import (
	"math"
	"sync"
)

type Comparison struct {
	sync.RWMutex

	base  float64
	price float64
}

func NewComparison() *Comparison {
	return &Comparison{
		base: math.NaN(),
	}
}

func (p *Comparison) Set(price float64) {
	p.Lock()
	defer p.Unlock()

	if math.IsNaN(p.base) {
		p.base = math.Log(price)
	}
	p.price = math.Log(price)
}

func (p *Comparison) Ratio() float64 {
	p.RLock()
	defer p.RUnlock()

	return math.Max(0, p.price/p.base)
}
