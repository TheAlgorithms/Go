package checksum

import "testing"

func TestCRC16(t *testing.T) {
    tests := []struct {
        input    string
        expected string
    }{
        {"hello", "A001"},
        {"world", "B5D0"},
        {"golang", "A34C"},
        {"", "A001"},
    }

    for _, tt := range tests {
        t.Run(tt.input, func(t *testing.T) {
            got := CRC16Hex(tt.input)
            if got != tt.expected {
                t.Errorf("CRC16Hex(%q) = %q; want %q", tt.input, got, tt.expected)
            }
        })
    }
}
