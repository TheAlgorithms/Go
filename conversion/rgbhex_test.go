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

func TestHEXToRGB(t *testing.T) {
	for i := 0; i < len(HEX); i++ {
		hex := HEX[i]
		expected := RGB[i]
		resultR, resultG, resultB := HEXToRGB(hex)
		if resultR != expected[0] || resultG != expected[1] || resultB != expected[2] {
			t.Errorf("HEXToRGB(%d) = %d,%d,%d; want %d,%d,%d",
				hex, resultR, resultG, resultB, expected[0], expected[1], expected[2])
		}
	}
}

func BenchmarkHEXToRGB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _, _ = HEXToRGB(0xdeadbe)
	}
}

func TestRGBToHEX(t *testing.T) {
	for i := 0; i < len(RGB); i++ {
		args := RGB[i]
		expected := HEX[i]
		result := RGBToHEX(args[0], args[1], args[2])
		if result != expected {
			t.Errorf("RGBToHEX(%d,%d,%d) = %d; want %d",
				args[0], args[1], args[2], result, expected)
		}
	}
}

func BenchmarkRGBToHEX(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = RGBToHEX(222, 173, 190)
	}
}
