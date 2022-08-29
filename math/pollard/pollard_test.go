// pollard_test.go
// description: Test for Pollard's rho algorithm
// author(s) [red_byte](https://github.com/i-redbyte)
// see pollard.go

package pollard

import (
	"math/big"
	"reflect"
	"testing"
)

func TestDefaultPolynomial(t *testing.T) {
	tests := []struct {
		name string
		x    *big.Int
		n    *big.Int
		want *big.Int
	}{
		{"Polynomial for n = 1772; x = 535", big.NewInt(535), big.NewInt(1772), big.NewInt(934)},
		{"Polynomial for n = 666; x = 53135", big.NewInt(53135), big.NewInt(666), big.NewInt(380)},
		{"Polynomial for n = 666; x = 13", big.NewInt(13), big.NewInt(666), big.NewInt(170)},
		{"Polynomial for n = 1917; x = 2510", big.NewInt(2510), big.NewInt(1917), big.NewInt(839)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := DefaultPolynomial(test.x, test.n); !reflect.DeepEqual(got, test.want) {
				t.Errorf("DefaultPolynomial() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestRho(t *testing.T) {
	tests := []struct {
		name    string
		n       *big.Int
		g       func(*big.Int, *big.Int) *big.Int
		want    *big.Int
		wantErr bool
	}{
		{"Rho for n:  11235 ", big.NewInt(11235), DefaultPolynomial, big.NewInt(21), false},
		{"Rho for n:  111155 ", big.NewInt(111155), DefaultPolynomial, big.NewInt(11), false},
		{"Rho for n:  8080 ", big.NewInt(8080), DefaultPolynomial, big.NewInt(16), false},
		{"Rho for n:  8536 ", big.NewInt(8536), DefaultPolynomial, big.NewInt(88), false},
		{"Rho for n:  666 ", big.NewInt(666), DefaultPolynomial, big.NewInt(3), false},
		{"Rho for n:  2 ", big.NewInt(2), DefaultPolynomial, big.NewInt(2), true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := Rho(test.n, test.g)
			if err != nil && !test.wantErr {
				t.Errorf("Rho() error = %v, wantErr %v", err, test.wantErr)
				return
			} else if err != nil && test.wantErr {
				return
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Rho() got = %v, want %v", got, test.want)
			}
		})
	}
}

func BenchmarkDefaultPolynomial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DefaultPolynomial(big.NewInt(535), big.NewInt(11235))
	}
}

func BenchmarkRho(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := Rho(big.NewInt(535), DefaultPolynomial)
		if err != nil {
			return
		}
	}
}
