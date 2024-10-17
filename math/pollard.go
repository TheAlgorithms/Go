// pollard.go
// description: Pollard's rho algorithm
// details:
// implementation of Pollard's rho algorithm for integer factorization-[Pollard's rho algorithm](https://en.wikipedia.org/wiki/Pollard%27s_rho_algorithm)
// time complexity: O(n^(1/4))
// space complexity: O(1)
// author(s) [red_byte](https://github.com/i-redbyte)
// see pollard_test.go

package math

import (
	"errors"
	"math/big"
)

// DefaultPolynomial is the commonly used polynomial g(x) = (x^2 + 1) mod n
func DefaultPolynomial(n *big.Int) func(*big.Int) *big.Int {
	bigOne := big.NewInt(1)
	bigTwo := big.NewInt(2)
	return func(x *big.Int) *big.Int {
		xSquared := new(big.Int).Exp(x, bigTwo, n) // see: https://en.wikipedia.org/wiki/Pollard%27s_rho_algorithm#Core_ideas
		xSquared.Add(xSquared, bigOne)
		xSquared.Mod(xSquared, n)
		return xSquared
	}
}

// PollardsRhoFactorization is an implementation of Pollard's rho factorization algorithm
// using the default parameters x = y = 2
func PollardsRhoFactorization(n *big.Int, f func(n *big.Int) func(x *big.Int) *big.Int) (*big.Int, error) {
	x, y, d := big.NewInt(2), big.NewInt(2), big.NewInt(1)
	bigOne := big.NewInt(1)
	g := f(n)
	for d.Cmp(bigOne) == 0 {
		x = g(x)
		y = g(g(y))
		sub := new(big.Int).Sub(x, y)
		d.GCD(nil, nil, sub.Abs(sub), n)
	}
	if d.Cmp(n) == 0 {
		return nil, errors.New("factorization failed")
	}
	return d, nil
}
