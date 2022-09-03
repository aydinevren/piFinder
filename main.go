package main

import (
	"fmt"
	"math/rand"
	"time"
)

func piFind(n int, threads int) (float64, time.Duration) {
	batchNumber := n / threads
	result := make(chan float64, threads)
	start := time.Now()

	for i := 0; i < threads; i++ {
		go func() {
			var sumCircle int
			r := rand.New(rand.NewSource(int64(rand.Float64())))

			for j := 0; j < batchNumber; j++ {
				x, y := r.Float64(), r.Float64()
				if x*x+y*y <= 1.0 {
					sumCircle++
				}
			}
			result <- 4.0 * float64(sumCircle) / float64(batchNumber)
		}()
	}
	var total float64
	for i := 0; i < threads; i++ {
		total += <-result
	}

	passingTime := time.Since(start)

	return total / float64(threads), passingTime
}

func main() {
	resultOfPi, passingTime := piFind(10000000000, 12)
	fmt.Println("Pi:", resultOfPi, "\nTime: ", passingTime)
}

/* func throwPebble() func() bool {
	return func() bool {
		r := rand.New(rand.NewSource(int64(rand.Float64())))
		x, y := r.Float64(), r.Float64()

		if x*x+y*y <= 1.0 {
			return true
		} else {
			return false
		}
	}
}

func piFind(n int, threads int) (float64, time.Duration) {
	batchNumber := n / threads
	result := make(chan float64, threads)
	start := time.Now()

	for i := 0; i < threads; i++ {
		go func() {
			var sumCircle int
			for j := 0; j < batchNumber; j++ {
				s := throwPebble()
				if s() {
					sumCircle++
				}
			}
			result <- 4.0 * float64(sumCircle) / float64(batchNumber)
		}()
	}
	var total float64
	for i := 0; i < threads; i++ {
		total += <-result
	}

	passingTime := time.Since(start)

	return total / float64(threads), passingTime
} */
