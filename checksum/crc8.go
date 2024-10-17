// crc8.go
// description: Calculate CRC8
// details:
// A cyclic redundancy check (CRC) is an error-detecting code commonly used in digital networks
// and storage devices to detect accidental changes to raw data.
// time complexity: O(n)
// space complexity: O(1)
// See more [CRC](https://en.wikipedia.org/wiki/Cyclic_redundancy_check)
// author(s) [red_byte](https://github.com/i-redbyte)
// see crc8_test.go

package checksum

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

// CRC8 calculates CRC8 checksum of the given data.
func CRC8(data []byte, model CRCModel) uint8 {
	table := getTable(model)
	crcResult := model.Init
	crcResult = addBytes(data, model, crcResult, table)
	if model.RefOut {
		crcResult = bits.Reverse8(crcResult)
	}
	return crcResult ^ model.XorOut
}

// This function get the result of adding the bytes in data to the crc
func addBytes(data []byte, model CRCModel, crcResult uint8, table []uint8) uint8 {
	if model.RefIn {
		for _, d := range data {
			d = bits.Reverse8(d)
			crcResult = table[crcResult^d]
		}
		return crcResult
	}
	for _, d := range data {
		crcResult = table[crcResult^d]
	}
	return crcResult
}

// This function get 256-byte (256x8) table for efficient processing.
func getTable(model CRCModel) []uint8 {
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
	return table
}
