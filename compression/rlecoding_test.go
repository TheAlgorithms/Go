package compression_test

import (
	"bytes"
	"testing"

	"github.com/TheAlgorithms/Go/compression"
)

func TestCompressionRLEncode(t *testing.T) {
	tests := []struct {
		name string
		data string
		want string
	}{
		{
			name: "test 1",
			data: "WWWWWWWWWWWWBWWWWWWWWWWWWBBB",
			want: "12W1B12W3B",
		},
		{
			name: "test 2",
			data: "AABCCCDEEEE",
			want: "2A1B3C1D4E",
		},
		{
			name: "test 3",
			data: "AAAABBBCCDA",
			want: "4A3B2C1D1A",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compression.RLEncode(tt.data); got != tt.want {
				t.Errorf("RLEncode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompressionRLEDecode(t *testing.T) {
	tests := []struct {
		name string
		data string
		want string
	}{
		{
			name: "test 1",
			data: "12W1B12W3B",
			want: "WWWWWWWWWWWWBWWWWWWWWWWWWBBB",
		},
		{
			name: "test 2",
			data: "2A1B3C1D4E",
			want: "AABCCCDEEEE",
		},
		{
			name: "test 3",
			data: "4A3B2C1D1A",
			want: "AAAABBBCCDA",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compression.RLEdecode(tt.data); got != tt.want {
				t.Errorf("RLEdecode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompressionRLEncodeBytes(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		want []byte
	}{
		{
			name: "test 1",
			data: []byte("WWWWWWWWWWWWBWWWWWWWWWWWWBBB"),
			want: []byte{12, 'W', 1, 'B', 12, 'W', 3, 'B'},
		},
		{
			name: "test 2",
			data: []byte("AABCCCDEEEE"),
			want: []byte{2, 'A', 1, 'B', 3, 'C', 1, 'D', 4, 'E'},
		},
		{
			name: "test 3",
			data: []byte("AAAABBBCCDA"),
			want: []byte{4, 'A', 3, 'B', 2, 'C', 1, 'D', 1, 'A'},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compression.RLEncodebytes(tt.data); !bytes.Equal(got, tt.want) {
				t.Errorf("RLEncodebytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompressionRLEDecodeBytes(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		want []byte
	}{
		{
			name: "test 1",
			data: []byte{12, 'W', 1, 'B', 12, 'W', 3, 'B'},
			want: []byte("WWWWWWWWWWWWBWWWWWWWWWWWWBBB"),
		},
		{
			name: "test 2",
			data: []byte{2, 'A', 1, 'B', 3, 'C', 1, 'D', 4, 'E'},
			want: []byte("AABCCCDEEEE"),
		},
		{
			name: "test 3",
			data: []byte{4, 'A', 3, 'B', 2, 'C', 1, 'D', 1, 'A'},
			want: []byte("AAAABBBCCDA"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compression.RLEdecodebytes(tt.data); !bytes.Equal(got, tt.want) {
				t.Errorf("RLEdecodebytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

/* --- BENCHMARKS --- */
func BenchmarkRLEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = compression.RLEncode("WWWWWWWWWWWWBWWWWWWWWWWWWBBB")
	}
}

func BenchmarkRLEDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = compression.RLEdecode("12W1B12W3B")
	}
}

func BenchmarkRLEncodeBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = compression.RLEncodebytes([]byte("WWWWWWWWWWWWBWWWWWWWWWWWWBBB"))
	}
}

func BenchmarkRLEDecodeBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = compression.RLEdecodebytes([]byte{12, 'W', 1, 'B', 12, 'W', 3, 'B'})
	}
}
