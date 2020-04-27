/*
	package prices

	Checksum: 指定価格帯注文の回避
*/

package price

import (
	"sync"
)

type Checker func(prefix interface{}, f float64) string
type Checksum struct {
	F      Checker
	prices sync.Map
}

func NewChecksum(f Checker) *Checksum {
	return &Checksum{
		F:      f,
		prices: sync.Map{},
	}
}

func (p *Checksum) IsOK(key interface{}) bool {
	_, isThere := p.prices.LoadOrStore(key, true)
	if isThere {
		// 同価格帯拒否
		return false
	}
	return true
}

func (p *Checksum) Close(key interface{}) {
	p.prices.Delete(key)
}
