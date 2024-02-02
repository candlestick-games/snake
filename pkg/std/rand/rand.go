package rand

import (
	"math"
	"math/rand"
)

func Float(min, max float64) float64 {
	if min == max {
		return min
	}
	if min > max {
		min, max = max, min
	}
	return rand.Float64()*(max-min) + min
}

func Int(min, max int) int {
	if min == max {
		return min
	}
	if min > max {
		min, max = max, min
	}
	return rand.Intn(max-min+1) + min
}

func IntWithSkew(min, max int, skew float64) int {
	if min == max {
		return min
	}
	if min > max {
		min, max = max, min
	}

	r := rand.Float64()
	return min + int(float64(max-min+1)*(1-math.Pow(r, skew)))
}

func Bool(probability float64) bool {
	if probability == 0 {
		return false
	}
	if probability >= 1 {
		return true
	}
	return rand.Float64() < probability
}
