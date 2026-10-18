package safe_sum

import "math"

// SafeSum последовательно суммирует транзакции, пропуская те, добавление
// которых привело бы к переполнению int64. skipped — индексы пропущенных
// элементов исходного среза, в порядке появления.
func SafeSum(transactions []int64) (total int64, skipped []int) {
	for i, t := range transactions {
		if t > 0 && total > math.MaxInt64-t {
			skipped = append(skipped, i)
			continue
		}
		if t < 0 && total < math.MinInt64-t {
			skipped = append(skipped, i)
			continue
		}
		total += t
	}
	return total, skipped
}
