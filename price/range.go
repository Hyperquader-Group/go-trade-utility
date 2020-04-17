package price

import (
	"math"
	"strings"
	"sync"
)

type Range struct {
	// digit 桁数
	digit int
	// n 桁番号
	n int

	prices sync.Map
}

// NewRange set digits and add price if same range price.
func NewRange(n, digit int) *Range {
	return &Range{
		digit:  digit,
		n:      n,
		prices: sync.Map{},
	}
}

// Set 設定桁数と桁番号をkeyにし、valueに注文価格を登録
func (p *Range) Set(sidein interface{}, price float64) (isThere bool) {
	var (
		side  = 1
		digit = float64(p.n) * math.Max(1, math.Pow10(p.digit))
	)
	switch v := sidein.(type) {
	case bool:
		if !v {
			side = -1
		}

	case int:
		side = v

	case string:
		if strings.ToLower(v) == "sell" {
			side = -1
		}
	}

	if side == 1 {
		_, isThere = p.prices.LoadOrStore(int(math.Floor(price/digit)*digit), price)
	} else if side == -1 {
		_, isThere = p.prices.LoadOrStore(int(math.Ceil(price/digit)*digit), price)
	}

	// すでにkey（価格帯に注文）があれば、true,なければfalse
	return isThere
}

func (p *Range) Reset() {
	p.prices = sync.Map{}
}
