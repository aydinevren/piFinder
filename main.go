package main

import (
	"fmt"
	"math/rand"
)

func throwPebble(r float64) func() bool {
	return func() bool {
		x := rand.Float64()*2 - 1
		y := rand.Float64()*2 - 1
		if x*x+y*y <= r {
			return true
		} else {
			return false
		}
	}
}

func piFind(n int) float64 {
	var sumCircle = 0
	var r = 1.0
	s := throwPebble(r)
	for i := 0; i < n; i++ {
		if s() {
			sumCircle++
		}
	}
	return 4.0 * float64(sumCircle) / float64(n)
}

func main() {
	fmt.Println(piFind(999))
}
