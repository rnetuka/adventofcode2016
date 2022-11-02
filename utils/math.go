package utils

import "math"

func Abs(n int) int {
	return int(math.Abs(float64(n)))
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
