package prime_test

import (
	"reflect"
	"testing"

	"github.com/TheAlgorithms/Go/math/prime"
)

func TestSieveEratosthenes(t *testing.T) {
	tests := []struct {
		name  string
		limit int
		want  []int
	}{
		{
			name:  "First 10 primes test",
			limit: 30,
			want:  []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29},
		},
		{
			name:  "First 20 primes test",
			limit: 71,
			want:  []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := prime.SieveEratosthenes(tt.limit)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SieveEratosthenes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSieveEratosthenes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = prime.SieveEratosthenes(10)
	}
}
