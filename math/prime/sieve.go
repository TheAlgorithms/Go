// sieve.go
// description: Algorithms for generating prime numbers efficiently
// author(s) [Taj](https://github.com/tjgurwara99)
// see sieve_test.go

package prime

// Generate generates the sequence of integers starting at 2 and sends it to the channel `ch`
func GenerateChannel(ch chan<- int) {
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

// Generate returns a int slice of prime numbers up to the limit
func Generate(limit int) []int {
	var primes []int

	ch := make(chan int)
	go GenerateChannel(ch)

	for i := 0; i < limit; i++ {
		primes = append(primes, <-ch)
		ch1 := make(chan int)
		go Sieve(ch, ch1, primes[i])
		ch = ch1
	}

	return primes
}
