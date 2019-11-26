package price

import (
	"gonum.org/v1/gonum/stat"
)

// DeviationRate 各取引所価格と比較したい場合など
func DeviationRate(price float64, prices, volumes []float64) (ratio, mean, stdv, zscore float64) {
	mean, stdv = stat.MeanStdDev(prices, volumes)
	zscore = stat.StdScore(price, mean, stdv)
	return price / mean, mean, stdv, zscore
}
