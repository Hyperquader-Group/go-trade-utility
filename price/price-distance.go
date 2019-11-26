package price

import "math"

// IsDistance 指定した価格差以上の指値とLTPの距離を検出
// 指値キャンセル時活用
func IsDistance(side int, ltp, orderPrice, distance float64) bool {
	if 0 < side {
		if ltp > orderPrice+math.Abs(distance*2) {
			return true
		}
		return false
	}

	if ltp < orderPrice-math.Abs(distance*2) {
		return true
	}
	return false
}

// IsExecuted 約定確認
// 同値の場合は未約定判定
func IsExecuted(orderside int, price, orderPrice float64) bool {
	if 0 < orderside { // 買い注文
		if price < orderPrice {
			return true
		}
	} else if orderside < 0 {
		if orderPrice < price {
			return true
		}
	}

	return false
}
