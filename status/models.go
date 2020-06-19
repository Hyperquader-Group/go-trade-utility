package status

import (
	"fmt"
	"math"
	"sync"
	"time"
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

	// 時間
	start time.Time

	// 注文と約定回数
	count      int
	executions int
}

func New() *OrderStates {
	return &OrderStates{

		start: time.Now(),
	}
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
	p.start = time.Now()
}

func (p *OrderStates) Rate() (countRate, sizeRate float64) {
	p.RLock()
	defer p.RUnlock()

	return math.Max(0, float64(p.executions)/float64(p.count)), math.Max(0, p.executionSize/p.orderSize)
}

func (p *OrderStates) String() string {
	p.RLock()
	defer p.RUnlock()

	return fmt.Sprintf(
		`注文回数: %d, 約定回数: %d, 約定率: %.2f％
注文枚数: %f, 約定枚数: %f, 約定率: %.2f％
集計開始: %s`,
		p.count, p.executions, math.Max(0, float64(p.executions)/float64(p.count)), *100,
		p.orderSize, p.executionSize, math.Max(0, p.executionSize/p.orderSize)*100,
		p.start.Format("2006/01/02 15:04:05"))
}
