// Package checksum provides functions for calculating various checksum values.
package checksum

import (
    "fmt"
)

// Adler32 computes the Adler-32 checksum of the input data.
func Adler32(data []byte) uint32 {
    var a, b uint32 = 1, 0
    for _, byte := range data {
        a = (a + uint32(byte)) % 65521
        b = (b + a) % 65521
    }
    return (b << 16) | a
}

// Adler32Hex computes the Adler-32 checksum of a string and returns it as a hexadecimal string.
func Adler32Hex(input string) string {
    adler := Adler32([]byte(input))
    return fmt.Sprintf("%08X", adler)
}
