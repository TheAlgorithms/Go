package conversion

import "testing"

var romanTestCases = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6,
	"VII": 7, "VIII": 8, "IX": 9, "X": 10, "XI": 11, "XII": 12,
	"XIII": 13, "XIV": 14, "XV": 15, "XVI": 16, "XVII": 17,
	"XVIII": 18, "XIX": 19, "XX": 20, "XXXI": 31, "XXXII": 32,
	"XXXIII": 33, "XXXIV": 34, "XXXV": 35, "XXXVI": 36, "XXXVII": 37,
	"XXXVIII": 38, "XXXIX": 39, "XL": 40, "XLI": 41, "XLII": 42,
	"XLIII": 43, "XLIV": 44, "XLV": 45, "XLVI": 46, "XLVII": 47,
	"XLVIII": 48, "XLIX": 49, "L": 50, "LXXXIX": 89, "XC": 90,
	"XCI": 91, "XCII": 92, "XCIII": 93, "XCIV": 94, "XCV": 95,
	"XCVI": 96, "XCVII": 97, "XCVIII": 98, "XCIX": 99, "C": 100,
	"CI": 101, "CII": 102, "CIII": 103, "CIV": 104, "CV": 105,
	"CVI": 106, "CVII": 107, "CVIII": 108, "CIX": 109, "CXLIX": 149,
	"CCCXLIX": 349, "CDLVI": 456, "D": 500, "DCIV": 604, "DCCLXXXIX": 789,
	"DCCCXLIX": 849, "CMIV": 904, "M": 1000, "MVII": 1007, "MLXVI": 1066,
	"MCCXXXIV": 1234, "MDCCLXXVI": 1776, "MMXXI": 2021, "MMDCCCVI": 2806,
	"MMCMXCIX": 2999, "MMM": 3000, "MMMCMLXXIX": 3979, "MMMCMXCIX": 3999,
}

func TestRomanToInt(t *testing.T) {
	for input, expected := range romanTestCases {
		out, err := RomanToInt(input)
		if err != nil {
			t.Errorf("RomanToInt(%s) returned an error %s", input, err.Error())
		}
		if out != expected {
			t.Errorf("RomanToInt(%s) = %d; want %d", input, out, expected)
		}
	}
	_, err := RomanToInt("IVCMXCIX")
	if err == nil {
		t.Error("RomanToInt(IVCMXCIX) expected an error")
	}

	val, err := RomanToInt("")
	if val != 0 {
		t.Errorf("RomanToInt(\"\") = %d; want 0", val)
	}
	if err != nil {
		t.Errorf("RomanToInt(\"\") returned an error %s", err.Error())
	}
}

func BenchmarkRomanToInt(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = RomanToInt("MMMCMXCIX")
	}
}
