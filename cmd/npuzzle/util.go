package main

import "math"

// (ceiling (log (max 1 num) 10)))
func NumWidth(n float64) uint {
	return uint(math.Ceil(math.Log10(n)))
}

func BoardNumWidth(board []int) uint {
	return NumWidth(float64(len(board)))
}

func GetY(n int, size int) int {
	if size < 0 {
		panic("Size < 0")
	}
	return n / size
}

func GetX(n int, size int) int {
	if size < 0 {
		panic("Size < 0")
	}
	return n % size
}
