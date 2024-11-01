package conversion

import (
	"testing"
)

func TestHexToDecimal(t *testing.T) {
	tests := []struct {
		hex     string
		want    int64
		wantErr bool
	}{
		{"", 0, true},
		{"G123", 0, true},
		{"123Z", 0, true},
		{"1", 1, false},
		{"A", 10, false},
		{"10", 16, false},
		{"1A", 26, false},
		{"aB", 171, false},
		{"0Ff", 255, false},
		{" 1A ", 26, false},
		{"0x1A", 26, false},
		{"0X1A", 26, false},
		{"1A", 26, false},
		{"7FFFFFFFFFFFFFFF", 9223372036854775807, false},
		{"0001A", 26, false},
		{"0000007F", 127, false},
		{"0", 0, false},
		{"0x0", 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.hex, func(t *testing.T) {
			got, err := hexToDecimal(tt.hex)
			if (err != nil) != tt.wantErr {
				t.Errorf("hexToDecimal(%q) error = %v, wantErr %v", tt.hex, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("hexToDecimal(%q) = %v, want %v", tt.hex, got, tt.want)
			}
		})
	}
}

func BenchmarkHexToDecimal(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = hexToDecimal("7FFFFFFFFFFFFFFF")
	}
}
