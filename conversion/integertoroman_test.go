package conversion

import "testing"

func TestIntToRoman(t *testing.T) {
	for expected, input := range romanTestCases {
		out, err := IntToRoman(input)
		if err != nil {
			t.Errorf("IntToRoman(%d) returned an error %s", input, err.Error())
		}
		if out != expected {
			t.Errorf("IntToRoman(%d) = %s; want %s", input, out, expected)
		}
	}
	_, err := IntToRoman(100000)
	if err == nil {
		t.Errorf("IntToRoman(%d) expected an error", 100000)
	}
	_, err = IntToRoman(0)
	if err == nil {
		t.Errorf("IntToRoman(%d) expected an error", 0)
	}
}

func BenchmarkIntToRoman(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = IntToRoman(3999)
	}
}
