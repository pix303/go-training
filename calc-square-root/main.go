package main

import (
	"flag"
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	var diff float64 = 2
	var prevDiff float64 = 0
	for math.Abs(diff - prevDiff) > 1e-15 {
		prevDiff = diff
		z -= (z*z - x) / (2*z)
		diff = math.Abs(z * z - x )
	}
	return z
}

func main() {
	input := flag.Float64("input",9, "input to calc square root")
	flag.Parse()

	fmt.Println(Sqrt(*input))
}
