package checksum

import "testing"

func TestAdler32(t *testing.T) {
    tests := []struct {
        input    string
        expected string
    }{
        {"hello", "1C2C8C41"},
        {"world", "1C0E26D2"},
        {"golang", "1C1B48C4"},
        {"", "00000001"},
    }

    for _, tt := range tests {
        t.Run(tt.input, func(t *testing.T) {
            got := Adler32Hex(tt.input)
            if got != tt.expected {
                t.Errorf("Adler32Hex(%q) = %q; want %q", tt.input, got, tt.expected)
            }
        })
    }
}
