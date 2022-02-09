// crc8_test.go
// description: Test for calculate crc8
// author(s) [red_byte](https://github.com/i-redbyte)
// see crc8.go

package crc

import (
	"fmt"
	"testing"
)

var (
	CRC8         = CRCModel{0x07, 0x00, false, false, 0x00, "CRC-8"}
	CRC8CDMA2000 = CRCModel{0x9B, 0xFF, false, false, 0x00, "CRC-8/CDMA2000"}
	CRC8DARC     = CRCModel{0x39, 0x00, true, true, 0x00, "CRC-8/DARC"}
	CRC8DVBS2    = CRCModel{0xD5, 0x00, false, false, 0x00, "CRC-8/DVB-S2"}
	CRC8EBU      = CRCModel{0x1D, 0xFF, true, true, 0x00, "CRC-8/EBU"}
	CRC8ICODE    = CRCModel{0x1D, 0xFD, false, false, 0x00, "CRC-8/I-CODE"}
	CRC8ITU      = CRCModel{0x07, 0x00, false, false, 0x55, "CRC-8/ITU"}
	CRC8MAXIM    = CRCModel{0x31, 0x00, true, true, 0x00, "CRC-8/MAXIM"}
	CRC8ROHC     = CRCModel{0x07, 0xFF, true, true, 0x00, "CRC-8/ROHC"}
	CRC8WCDMA    = CRCModel{0x9B, 0x00, true, true, 0x00, "CRC-8/WCDMA"}
)

func TestCalculateCRC8(t *testing.T) {
	data := []byte("Hello World")
	tests := []struct {
		name  string
		data  []byte
		model CRCModel
		want  uint8
	}{
		{CRC8.Name, data, CRC8, 0x25},
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
			fmt.Println(CalculateCRC8(test.data, test.model))
			if got := CalculateCRC8(test.data, test.model); got != test.want {
				t.Errorf("CalculateCRC8() = %v, want %v", got, test.want)
			}
		})
	}
}

func BenchmarkCalculateCRC8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateCRC8([]byte("123456789"), CRC8)
	}
}
