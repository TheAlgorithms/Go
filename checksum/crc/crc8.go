// crc8.go
// description: Calculate CRC8
// details:
// A cyclic redundancy check (CRC) is an error-detecting code commonly used in digital networks and storage devices to detect accidental changes to raw data. See more [CRC](https://en.wikipedia.org/wiki/Cyclic_redundancy_check)
// author(s) [red_byte](https://github.com/i-redbyte)
// see crc8_test.go

package crc

import "math/bits"

// CRCModel contains the necessary parameters for calculating the DRC algorithm
type CRCModel struct {
	Poly   uint8
	Init   uint8
	RefIn  bool
	RefOut bool
	XorOut uint8
	Name   string
}

// CalculateCRC8 This function calculate CRC8 checksum.
func CalculateCRC8(data []byte, model CRCModel) uint8 {
	table := make([]uint8, 256)
	for i := 0; i < 256; i++ {
		crc := uint8(i)
		for j := 0; j < 8; j++ {
			isSetBit := (crc & 0x80) != 0
			crc <<= 1
			if isSetBit {
				crc ^= model.Poly
			}
		}
		table[i] = crc
	}
	crcResult := model.Init

	if model.RefIn {
		for _, d := range data {
			d = bits.Reverse8(d)
			crcResult = table[crcResult^d]
		}
	} else {
		for _, d := range data {
			crcResult = table[crcResult^d]
		}
	}
	if model.RefOut {
		crcResult = bits.Reverse8(crcResult)
	}

	return crcResult ^ model.XorOut
}
