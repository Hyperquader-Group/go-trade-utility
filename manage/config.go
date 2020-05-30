package manage

import "time"

type Config struct {
	ProductCode, SubProductCode string
	// 集計期間
	SetTerm  int
	TermUnit time.Duration

	// ロット
	Lots    []float64
	LotMinN int
	LotMaxN int
}

func NewConfig(term int, productcode string, usesizes []float64) *Config {
	return &Config{
		ProductCode: productcode,
		SetTerm:     term,
		TermUnit:    time.Second,

		Lots:    usesizes,
		LotMinN: 0,
		LotMaxN: 1,
	}
}

func (p *Config) Term() time.Duration {
	return time.Duration(p.SetTerm) * p.TermUnit
}
