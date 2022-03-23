// crc8_test.go
// description: Test for calculate crc8
// author(s) [red_byte](https://github.com/i-redbyte)
// see crc8.go

package checksum_test

import (
	"fmt"
	"github.com/TheAlgorithms/Go/checksum"
	"testing"
)

func TestCRC8(t *testing.T) {
	data := []byte("123456789")
	tests := []struct {
		name  string
		data  []byte
		model checksum.CRCModel
		want  uint8
	}{
		{CRC8.Name, data, CRC8, 0xF4},
		{CRC8CDMA2000.Name, data, CRC8CDMA2000, 0xDA},
		{CRC8DARC.Name, data, CRC8DARC, 0x15},
		{CRC8DVBS2.Name, data, CRC8DVBS2, 0xBC},
		{CRC8EBU.Name, data, CRC8EBU, 0x97},
		{CRC8ICODE.Name, data, CRC8ICODE, 0x7E},
		{CRC8ITU.Name, data, CRC8ITU, 0xA1},
		{CRC8MAXIM.Name, data, CRC8MAXIM, 0xA1},
		{CRC8ROHC.Name, data, CRC8ROHC, 0xD0},
		{CRC8WCDMA.Name, data, CRC8WCDMA, 0x25},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			fmt.Println(checksum.CRC8(test.data, test.model))
			if got := checksum.CRC8(test.data, test.model); got != test.want {
				t.Errorf("CRC8() = %v, want %v", got, test.want)
			}
		})
	}
}

func BenchmarkCRC8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checksum.CRC8([]byte("123456789"), CRC8)
	}
}

var (
	CRC8 = checksum.CRCModel{
		Poly:   0x07,
		Init:   0x00,
		RefIn:  false,
		RefOut: false,
		XorOut: 0x00,
		Name:   "CRC-8",
	}
	CRC8CDMA2000 = checksum.CRCModel{
		Poly:   0x9B,
		Init:   0xFF,
		RefIn:  false,
		RefOut: false,
		XorOut: 0x00,
		Name:   "CRC-8/CDMA2000",
	}
	CRC8DARC = checksum.CRCModel{
		Poly:   0x39,
		Init:   0x00,
		RefIn:  true,
		RefOut: true,
		XorOut: 0x00,
		Name:   "CRC-8/DARC",
	}
	CRC8DVBS2 = checksum.CRCModel{
		Poly:   0xD5,
		Init:   0x00,
		RefIn:  false,
		RefOut: false,
		XorOut: 0x00,
		Name:   "CRC-8/DVB-S2",
	}
	CRC8EBU = checksum.CRCModel{
		Poly:   0x1D,
		Init:   0xFF,
		RefIn:  true,
		RefOut: true,
		XorOut: 0x00,
		Name:   "CRC-8/EBU",
	}
	CRC8ICODE = checksum.CRCModel{
		Poly:   0x1D,
		Init:   0xFD,
		RefIn:  false,
		RefOut: false,
		XorOut: 0x00,
		Name:   "CRC-8/I-CODE",
	}
	CRC8ITU = checksum.CRCModel{
		Poly:   0x07,
		Init:   0x00,
		RefIn:  false,
		RefOut: false,
		XorOut: 0x55,
		Name:   "CRC-8/ITU",
	}
	CRC8MAXIM = checksum.CRCModel{
		Poly:   0x31,
		Init:   0x00,
		RefIn:  true,
		RefOut: true,
		XorOut: 0x00,
		Name:   "CRC-8/MAXIM",
	}
	CRC8ROHC = checksum.CRCModel{
		Poly:   0x07,
		Init:   0xFF,
		RefIn:  true,
		RefOut: true,
		XorOut: 0x00,
		Name:   "CRC-8/ROHC",
	}
	CRC8WCDMA = checksum.CRCModel{
		Poly:   0x9B,
		Init:   0x00,
		RefIn:  true,
		RefOut: true,
		XorOut: 0x00,
		Name:   "CRC-8/WCDMA",
	}
)
