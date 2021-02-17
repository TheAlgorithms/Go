package sieve

// Generate generates the sequence of integers starting at 2 and sends it to the channel `ch`
func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

// Sieve Sieving the numbers that are not prime from the channel - basically removing them from the channels
func Sieve(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}
