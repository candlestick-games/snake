package rand

import "math/rand"

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
