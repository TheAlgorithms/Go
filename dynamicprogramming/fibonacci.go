package dynamicprogramming

// Nth fibonacci Number
// https://www.geeksforgeeks.org/program-for-nth-fibonacci-number/

func fib(n int) int {
	/* Declare an array to store Fibonacci numbers. */
	f := make([]int, n+2) // 1 extra to handle case, n = 0

	/* 0th and 1st number of the series are 0 and 1*/
	f[0] = 0
	f[1] = 1

	for i := 2; i <= n; i++ {
		/* Add the previous 2 numbers in the series
		and store it */
		f[i] = f[i-1] + f[i-2]
	}

	return f[n]
}
