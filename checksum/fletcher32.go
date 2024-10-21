// Package checksum provides functions for calculating various checksum values.
package checksum

import (
    "fmt"
)

// Fletcher32 computes the Fletcher-32 checksum of the input data.
func Fletcher32(data []byte) uint32 {
    var sum1, sum2 uint32 = 0xff, 0xff
    for _, byte := range data {
        sum1 = (sum1 + uint32(byte)) % 0xff
        sum2 = (sum2 + sum1) % 0xff
    }
    return (sum2 << 8) | sum1
}

// Fletcher32Hex computes the Fletcher-32 checksum of a string and returns it as a hexadecimal string.
func Fletcher32Hex(input string) string {
    fletcher := Fletcher32([]byte(input))
    return fmt.Sprintf("%08X", fletcher)
}
