package scientificcalculator

import (
	"github.com/sbrunettajr/calculator"
)

func Abs(v float64) float64 {
	if v < 0 {
		return calculator.Mul(v, -1)
	}
	return v
}
