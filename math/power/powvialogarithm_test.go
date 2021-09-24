// powvialogarithm_test.go
// description: Test for UsingLog
// author(s) [red_byte](https://github.com/i-redbyte)
// see powvialogarithm.go

package power

import "testing"

func TestUsingLog(t *testing.T) {
	var tests = []struct {
		name     string
		base     float64
		power    float64
		expected float64
	}{
		{"0^0", 99, 1, 99},
		{"-3^9", -3, 9, -19683},
		{"0^2", 0, 2, 0},
		{"2^0", 2, 0, 1},
		{"2^3", 2, 3, 8},
		{"8^3", 8, 3, 512},
		{"11^11", 11, 11, 285311670611},
		{"5^5", 5, 5, 3125},
		{"-7^2", -7, 2, 49},
		{"-6^3", -6, 3, -216},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := UsingLog(tc.base, tc.power)
			t.Log(result)
			if result != tc.expected {
				t.Errorf("Expected %.2f to the power of %.2f to be: %.2f, but got: %.2f", tc.base, tc.power, tc.expected, result)
			}
		})
	}
}

func BenchmarkUsingLog(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UsingLog(10, 5)
	}
}
