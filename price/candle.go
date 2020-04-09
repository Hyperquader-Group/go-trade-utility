package price

import (
	"fmt"
	"math"
	"sync"
)

type OHLCver interface {
	Hige() (isUpward, tooNarrow bool, price, hige, volume float64)
}

type OHLCv struct {
	Open, Close float64
	High, Low   float64
	Volume      float64
}

type Candle struct {
	mux sync.RWMutex

	open   float64
	high   float64
	low    float64
	close  float64
	volume float64

	prices  []float64
	volumes []float64
}

func NewCandle() *Candle {
	return &Candle{
		prices:  make([]float64, 0),
		volumes: make([]float64, 0),
	}
}

func (p *Candle) Set(price, volume float64) {
	p.mux.Lock()
	defer p.mux.Unlock()

	if len(p.prices) == 0 {
		p.open = price
		p.high = math.Inf(-1)
		p.low = math.Inf(1)
		p.volume = 0
	}
	p.prices = append(p.prices, price)
	p.volumes = append(p.volumes, volume)
	p.volume += volume

	if p.high < price {
		p.high = price
	}
	if p.low > price {
		p.low = price
	}

}

func (p *Candle) Reset() OHLCv {
	p.mux.Lock()
	defer p.mux.Unlock()

	if 0 < len(p.prices) {
		p.close = p.prices[len(p.prices)-1]
	} else {
		p.close = p.open
	}
	p.prices = make([]float64, 0)
	p.volumes = make([]float64, 0)

	return OHLCv{
		Open:   p.open,
		High:   p.high,
		Low:    p.low,
		Close:  p.close,
		Volume: p.volume,
	}
}

func (p *Candle) Open() float64 {
	p.mux.RLock()
	defer p.mux.RUnlock()
	return p.open
}

func (p *Candle) Close() float64 {
	p.mux.RLock()
	defer p.mux.RUnlock()
	return p.close
}

func (p *Candle) High() float64 {
	p.mux.RLock()
	defer p.mux.RUnlock()
	return p.high
}

func (p *Candle) Low() float64 {
	p.mux.RLock()
	defer p.mux.RUnlock()
	return p.low
}

func (p *Candle) Volume() float64 {
	p.mux.RLock()
	defer p.mux.RUnlock()
	return p.volume
}

type HigeType int

func (p HigeType) String() string {
	switch p {
	case Cross:
		return "Cross body"
	case Hummer:
		return "Hummer body"

	}
	return "Box body"
}

const (
	_ HigeType = iota
	Cross
	BoxBody
	Hummer
)

type Status int

func (p Status) String() string {
	switch p {
	case Positive:
		return "Positive"
	case Negative:
		return "Negative"

	}
	return "Flat"
}

const (
	Flat Status = iota
	Positive
	Negative
)

type HIGEvaluate struct {
	Type         HigeType
	IsSufficient bool

	Overall float64

	Body       float64
	BodyStatus Status

	Hige       float64
	HigeStatus Status

	Top, Bottom float64
}

func (p *Candle) Hige(changeRate float64) HIGEvaluate {
	p.mux.RLock()
	defer p.mux.RUnlock()

	var hige HIGEvaluate

	// BodyBoxを定義する
	hige.Body = p.open - p.close
	if 0 < hige.Body {
		hige.BodyStatus = Negative
	} else if hige.Body < 0 {
		hige.BodyStatus = Positive
	} else {
		hige.BodyStatus = Flat
	}
	bodySize := math.Abs(hige.Body)
	hige.Overall = math.Abs(p.high - p.low)
	hige.Body = math.Max(0, bodySize/hige.Overall)

	// ひげの方向と長さを取得
	var (
		upper, lower float64
	)
	if p.open < p.close {
		upper = math.Max(0, p.high-p.close)
		lower = math.Max(0, p.open-p.low)
	} else if p.open > p.close {
		upper = math.Max(0, p.high-p.open)
		lower = math.Max(0, p.close-p.low)
	}

	if upper < lower {
		// 下ヒゲの比率
		hige.Hige = math.Max(0, lower/hige.Overall)
		hige.HigeStatus = Negative
	} else if upper > lower {
		// 上ヒゲの比率
		hige.Hige = math.Max(0, upper/hige.Overall)
		hige.HigeStatus = Positive
	} else {
		hige.HigeStatus = Flat
	}

	// 特徴的な足型の定義
	if 0.80 < hige.Hige {
		hige.Type = Hummer
	} else if 0.8 < hige.Body {
		hige.Type = BoxBody
	} else if hige.Body < 0.2 {
		hige.Type = Cross
	}

	// 上限下限
	hige.Top = p.high
	hige.Bottom = p.low

	// changeRate以上の価格の変動を分別する
	threshold := p.open * math.Abs(changeRate)
	if hige.Overall < threshold {
		return hige
	}

	hige.IsSufficient = true
	return hige
}

func (p HIGEvaluate) String() string {
	if p.Type == Hummer && p.BodyStatus == Positive && p.HigeStatus == Negative {
		return fmt.Sprintf("買い注文 type:%s body:%s hige:%s", p.Type, p.BodyStatus, p.HigeStatus)
	} else if p.Type == Hummer && p.BodyStatus == Negative && p.HigeStatus == Positive {
		return fmt.Sprintf("売り注文 type:%s body:%s hige:%s", p.Type, p.BodyStatus, p.HigeStatus)
	} else if p.Type == Hummer && p.BodyStatus == Positive && p.HigeStatus == Positive {
		return fmt.Sprintf("買い検討 type:%s body:%s hige:%s", p.Type, p.BodyStatus, p.HigeStatus)
	} else if p.Type == Hummer && p.BodyStatus == Negative && p.HigeStatus == Negative {
		return fmt.Sprintf("売り検討 type:%s body:%s hige:%s", p.Type, p.BodyStatus, p.HigeStatus)
	}
	return fmt.Sprintf("待機 type:%s body:%s hige:%s", p.Type, p.BodyStatus, p.HigeStatus)
}
