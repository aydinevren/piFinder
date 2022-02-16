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

func piFind(n int, threads int) float64 {
	threadsNumber := n / threads
	result := make(chan float64, threads)
	for i := 0; i < threads; i++ {
		go func() {
			var sumCircle = 0
			var r = 1.0
			s := throwPebble(r)
			for j := 0; j < threadsNumber; j++ {
				if s() {
					sumCircle++
				}
			}
			result <- 4.0 * float64(sumCircle) / float64(threadsNumber)

		}()
	}
	var total float64
	for i := 0; i < threads; i++ {
		total += <-result
	}

	return total / float64(threads)
}

func main() {
	fmt.Println(piFind(1000000000, 2))
}
