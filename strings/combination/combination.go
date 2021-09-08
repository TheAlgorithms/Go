// Package combination ...
package combination

import "fmt"

// Combinations structure with in and out rune
type Combinations struct {
	out []rune
	in  []rune
}

// Start ...
func Start(input string) {
	c := &Combinations{
		in: []rune(input),
	}

	c.Combine(0)
}

// Combine ...
func (c *Combinations) Combine(seed int) {
	inLen := len(c.in)
	for i := seed; i < inLen-1; i++ {
		c.out = append(c.out, c.in[i])
		fmt.Println(string(c.out))
		c.Combine(i + 1)
		c.out = c.out[:len(c.out)-1]
	}
	c.out = append(c.out, c.in[inLen-1])
	fmt.Println(string(c.out))
	c.out = c.out[:len(c.out)-1]
}
