package price

import (
	"time"

	"gonum.org/v1/gonum/stat/distuv"
)

// Poissons 設定時間単位の出現を保持し、もっとも出現しやすい出現価格を返す
type Poissons struct {
	Term    time.Duration
	Appears []float64

	Po *distuv.Poisson
}
