package status

import (
	"fmt"
	"math"
	"sync"
)

const (
	ORDER = iota
	ORDEREXECUTION
)

type OrderStates struct {
	sync.RWMutex

	// 注文枚数
	orderSize     float64
	executionSize float64

	// 注文と約定回数
	count      int
	executions int
}

func New() *OrderStates {
	return &OrderStates{}
}

func (p *OrderStates) Set(t int, f float64) {
	p.Lock()
	defer p.Unlock()

	switch t {
	case ORDER: // 注文
		p.orderSize += f
		p.count++

	case ORDEREXECUTION: // 約定
		p.executionSize += f
		p.executions++
	}

}

func (p *OrderStates) Reset() {
	p.Lock()
	defer p.Unlock()

	p.orderSize = 0
	p.count = 0
	p.executionSize = 0
	p.executions = 0
}

func (p *OrderStates) Rate() (countRate, sizeRate float64) {
	p.RLock()
	defer p.RUnlock()

	return math.Max(0, p.orderSize/p.executionSize), math.Max(0, float64(p.count)/float64(p.executions))
}

func (p *OrderStates) String() string {
	p.RLock()
	defer p.RUnlock()

	return fmt.Sprintf(
		`注文回数: %d, 約定回数: %d, 約定率: %.2f％
注文枚数: %f, 約定枚数: %f, 約定率: %.2f％`,
		p.count, p.executions, math.Max(0, p.orderSize/p.executionSize)*100,
		p.orderSize, p.executionSize, math.Max(0, float64(p.count)/float64(p.executions))*100)
}
