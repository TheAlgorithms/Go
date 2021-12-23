package binary

import "testing"

func TestAbs(t *testing.T) {
	tests := getAbsTests()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Abs(test.base, test.n); got != test.want {
				t.Errorf("Abs() = %v, want %v", got, test.want)
			}
		})
	}
}

func getAbsTests() []struct {
	name string
	base int
	n    int
	want int
} {
	tests := []struct {
		name string
		base int
		n    int
		want int
	}{
		{"-1 = |1| ", 32, -1, 1},
		{"-255 = |255| ", 32, -255, 255},
		{"0 = |0| ", 64, 0, 0},
		{"5 = |5| ", 16, 5, 5},
		{"-5 = |5| ", 32, -5, 5},
		{"-98368972 = |98368972| ", 64, -98368972, 98368972},
		{"-110298368972 = |110298368972| ", 64, -98368972, 98368972},
	}
	return tests
}
