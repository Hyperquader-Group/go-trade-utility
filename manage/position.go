package manage

import "math"

// Position is Positions struct
type Position struct {
	Min   float64
	Size  float64
	Limit float64
}

// NewPosition is new Positions struct
func NewPosition(min, limit float64) *Position {
	return &Position{
		Min:   min,
		Limit: limit,
	}
}

// Set is sets size
func (p *Position) Set(size float64) {
	p.Size = size
}

// Lot is Size for order lot
// size, sideに対応する新規枚数が返ってくる
func (p *Position) Lot(side int, tension float64) (bool, float64) {
	bias := p.bias(tension)
	size := p.Limit * bias
	if p.isFull(side, size) {
		return true, 0
	}

	size = p.checkSame(side, size)

	return false, math.Max(p.Min, math.Abs(size))
}

// bias is positions bias
func (p *Position) bias(tension float64) float64 {
	return math.Tanh(tension * p.Size / p.Limit)
}

// isFull is checks Limit&Size
// 売り方向要望を受け、同方向建玉過多ならばisFull
func (p *Position) isFull(side int, size float64) bool {
	if 0 < side { // 新規買い注文
		if p.Limit < p.Size {
			return true
		}
		if p.Limit < p.Size+math.Abs(size) {
			return true
		}
		if p.Limit < p.Size+p.Min {
			return true
		}

	} else if side < 0 { // 新規売り注文
		if p.Size < -p.Limit {
			return true
		}
		if p.Size-math.Abs(size) < -p.Limit {
			return true
		}
		if p.Size-math.Abs(p.Min) < -p.Limit {
			return true
		}

	}

	return false
}

// checkSame 注文多重化の際、買い建玉に買い注文過多を避ける目的
func (p *Position) checkSame(side int, size float64) float64 {
	if p.Min < math.Abs(size) {
		if 0 < side && 0 < p.Size {
			size = p.Min
		} else if side < 0 && p.Size < 0 {
			size = p.Min
		}
	}
	return size
}
