package main

import "math"

func NumWidth(n float64) uint {
	return uint(math.Ceil(math.Log10(n)))
}


func BoardNumWidth(board []uint) uint {
	var max uint
	max = 0
	for _, n := range board {
		if n > max {
			max = n
		}
	}
	width := NumWidth(float64(max))
	return width
}
