package checksum

import "testing"

func TestFletcher32(t *testing.T) {
    tests := []struct {
        input    string
        expected string
    }{
        {"hello", "5E3C5E27"},
        {"world", "5E205B9C"},
        {"golang", "5E1F73CC"},
        {"", "000000FF"},
    }

    for _, tt := range tests {
        t.Run(tt.input, func(t *testing.T) {
            got := Fletcher32Hex(tt.input)
            if got != tt.expected {
                t.Errorf("Fletcher32Hex(%q) = %q; want %q", tt.input, got, tt.expected)
            }
        })
    }
}
