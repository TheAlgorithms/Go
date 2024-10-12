// Package checksum provides functions for calculating various checksum values.
package checksum

import "fmt"

// CRC16 computes the CRC-16 checksum of the input data.
func CRC16(data []byte) uint16 {
    var crc uint16 = 0xFFFF
    for _, b := range data {
        crc ^= uint16(b)
        for i := 0; i < 8; i++ {
            if crc&0x0001 != 0 {
                crc = (crc >> 1) ^ 0xA001
            } else {
                crc >>= 1
            }
        }
    }
    return crc
}

// CRC16Hex computes the CRC-16 checksum of a string and returns it as a hexadecimal string.
func CRC16Hex(input string) string {
    crc := CRC16([]byte(input))
    return fmt.Sprintf("%04X", crc)
}
