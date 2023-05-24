package conversion

import "testing"

func TestIntegerToRoman(t *testing.T) {
	for expected, input := range romanTestCases {
		out, err := IntegerToRoman(input)
		if err != nil {
			t.Errorf("IntegerToRoman(%d) returned an error %s", input, err.Error())
		}
		if out != expected {
			t.Errorf("IntegerToRoman(%d) = %s; want %s", input, out, expected)
		}
	}
	_, err := IntegerToRoman(100000)
	if err == nil {
		t.Errorf("IntegerToRoman(%d) expected an error", 100000)
	}
	_, err = IntegerToRoman(0)
	if err == nil {
		t.Errorf("IntegerToRoman(%d) expected an error", 0)
	}
}

func BenchmarkIntegerToRoman(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = IntegerToRoman(3999)
	}
}
