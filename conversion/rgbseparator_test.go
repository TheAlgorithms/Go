package conversion

import "testing"

var HEX = []uint{
	0x1abc9c,
	0x3498db,
	0x9b59b6,
}

var RGB = [][]byte{
	{26, 188, 156},
	{52, 152, 219},
	{155, 89, 182},
}

func TestExtractRGB(t *testing.T) {
	for i := 0; i < len(HEX); i++ {
		hex := HEX[i]
		expected := RGB[i]
		resultR, resultG, resultB := ExtractRGB(hex)
		if resultR != expected[0] || resultG != expected[1] || resultB != expected[2] {
			t.Errorf("ExtractRGB(%d) = %d,%d,%d; want %d,%d,%d",
				hex, resultR, resultG, resultB, expected[0], expected[1], expected[2])
		}
	}
}

func BenchmarkExtractRGB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _, _ = ExtractRGB(0xdeadbe)
	}
}

func TestCombineRGB(t *testing.T) {
	for i := 0; i < len(RGB); i++ {
		args := RGB[i]
		expected := HEX[i]
		result := CombineRGB(args[0], args[1], args[2])
		if result != expected {
			t.Errorf("CombineRGB(%d,%d,%d) = %d; want %d",
				args[0], args[1], args[2], result, expected)
		}
	}
}

func BenchmarkCombineRGB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = CombineRGB(222, 173, 190)
	}
}
