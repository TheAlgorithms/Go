package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	cpus := runtime.NumCPU()
	monteCarlo := func(iterations, cpus int) float64{
		start := time.Now()
		flow := make(chan int, int(math.Floor(float64(cpus)/2)))
		var wg sync.WaitGroup
		var total int
		miniFunc := func(iters int, flow chan int) {
			defer wg.Done()
			var inside int
			for i := 0; i < iters; i++ {
				x := rand.Float64()
				y := rand.Float64()
				if math.Hypot(x, y) <= 1 {
					inside++
				}
			}
			flow <- inside
		} //end of minifunc
		fuxc := func(flow chan int) {
			defer wg.Done()
			for i := 0; i < int(math.Floor(float64(cpus)/2)); i++ {
				total += <-flow
			}
		} //end of fuxc
		wg.Add(int(math.Floor(float64(cpus) / 2)))
		//spawn goroutines for half the logical processors
		// This leaves free cpus for swapping
		for i := 0; i < int(math.Floor(float64(cpus)/2)); i++ {
			go miniFunc(int(math.Ceil(float64(iterations/int(math.Floor(float64(cpus)/2))))), flow)
		}
		wg.Add(1)
		go fuxc(flow)
		wg.Wait()
		fmt.Printf("took %v \n",time.Since(start))
		return 4 * float64(total) / float64(iterations)
	}//end of monteCarlo
	fmt.Println(monteCarlo(2<<18, cpus))
}
