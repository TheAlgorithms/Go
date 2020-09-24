package main

import (
	_ "fmt"
)

func main() {
	var l, f int
	l = 100
	channel := make(chan int, l)
	channelTwo := make(chan int, l)
	var nums []bool

	for i := 0; i < l; i++ {
		nums = append(nums, true)
	}
	f = 2

	go crossOff(channel, channelTwo, f, l)
	go nextF(f, channelTwo, nums)
	//blocks and sieves out non-primes
	for {
		select {
		case index := <-channel:
			nums[index-1] = false
		}
	}
}
func crossOff(c chan int, r chan int, f int, l int) {
	for f != -1 {
		for f <= l {
			c <- f
			f += f
		}
		f = <-r
		r <- 1
	}

}

func nextF(f int, r chan int, nums []bool) {
	for {
		<-r
		for i, v := range nums[f:] {
			if v == true {
				r <- i + f
			}
		}
	}
}
