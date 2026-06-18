package conversion

import "testing"

func TestBase64Encode(t *testing.T) {
	testCases := []struct {
		in       string
		expected string
	}{
		{"Hello World!", "SGVsbG8gV29ybGQh"},       // multiple of 3 byte length (multiple of 24-bits)
		{"Hello World!a", "SGVsbG8gV29ybGQhYQ=="},  // multiple of 3 byte length + 1
		{"Hello World!ab", "SGVsbG8gV29ybGQhYWI="}, // multiple of 3 byte length + 2
		{"", ""},      // empty byte slice
		{"6", "Ng=="}, // short text
		{"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.", "TG9yZW0gaXBzdW0gZG9sb3Igc2l0IGFtZXQsIGNvbnNlY3RldHVyIGFkaXBpc2NpbmcgZWxpdCwgc2VkIGRvIGVpdXNtb2QgdGVtcG9yIGluY2lkaWR1bnQgdXQgbGFib3JlIGV0IGRvbG9yZSBtYWduYSBhbGlxdWEuIFV0IGVuaW0gYWQgbWluaW0gdmVuaWFtLCBxdWlzIG5vc3RydWQgZXhlcmNpdGF0aW9uIHVsbGFtY28gbGFib3JpcyBuaXNpIHV0IGFsaXF1aXAgZXggZWEgY29tbW9kbyBjb25zZXF1YXQuIER1aXMgYXV0ZSBpcnVyZSBkb2xvciBpbiByZXByZWhlbmRlcml0IGluIHZvbHVwdGF0ZSB2ZWxpdCBlc3NlIGNpbGx1bSBkb2xvcmUgZXUgZnVnaWF0IG51bGxhIHBhcmlhdHVyLiBFeGNlcHRldXIgc2ludCBvY2NhZWNhdCBjdXBpZGF0YXQgbm9uIHByb2lkZW50LCBzdW50IGluIGN1bHBhIHF1aSBvZmZpY2lhIGRlc2VydW50IG1vbGxpdCBhbmltIGlkIGVzdCBsYWJvcnVtLg=="}, // Long text
	}

	for _, tc := range testCases {
		result := Base64Encode([]byte(tc.in))
		if result != tc.expected {
			t.Fatalf("Base64Encode(%s) = %s, want %s", tc.in, result, tc.expected)
		}
	}
}

func BenchmarkBase64Encode(b *testing.B) {
	benchmarks := []struct {
		name     string
		in       string
		expected string
	}{
		{"Hello World!", "Hello World!", "SGVsbG8gV29ybGQh"},         // multiple of 3 byte length (multiple of 24-bits)
		{"Hello World!a", "Hello World!a", "SGVsbG8gV29ybGQhYQ=="},   // multiple of 3 byte length + 1
		{"Hello World!ab", "Hello World!ab", "SGVsbG8gV29ybGQhYWI="}, // multiple of 3 byte length + 2
		{"Empty", "", ""},  // empty byte slice
		{"6", "6", "Ng=="}, // short text
		{"Lorem ipsum", "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.", "TG9yZW0gaXBzdW0gZG9sb3Igc2l0IGFtZXQsIGNvbnNlY3RldHVyIGFkaXBpc2NpbmcgZWxpdCwgc2VkIGRvIGVpdXNtb2QgdGVtcG9yIGluY2lkaWR1bnQgdXQgbGFib3JlIGV0IGRvbG9yZSBtYWduYSBhbGlxdWEuIFV0IGVuaW0gYWQgbWluaW0gdmVuaWFtLCBxdWlzIG5vc3RydWQgZXhlcmNpdGF0aW9uIHVsbGFtY28gbGFib3JpcyBuaXNpIHV0IGFsaXF1aXAgZXggZWEgY29tbW9kbyBjb25zZXF1YXQuIER1aXMgYXV0ZSBpcnVyZSBkb2xvciBpbiByZXByZWhlbmRlcml0IGluIHZvbHVwdGF0ZSB2ZWxpdCBlc3NlIGNpbGx1bSBkb2xvcmUgZXUgZnVnaWF0IG51bGxhIHBhcmlhdHVyLiBFeGNlcHRldXIgc2ludCBvY2NhZWNhdCBjdXBpZGF0YXQgbm9uIHByb2lkZW50LCBzdW50IGluIGN1bHBhIHF1aSBvZmZpY2lhIGRlc2VydW50IG1vbGxpdCBhbmltIGlkIGVzdCBsYWJvcnVtLg=="}, // Long text
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Base64Encode([]byte(bm.in))
			}
		})
	}
}

func TestBase64Decode(t *testing.T) {
	testCases := []struct {
		expected string
		in       string
	}{
		{"Hello World!", "SGVsbG8gV29ybGQh"},       // multiple of 3 byte length (multiple of 24-bits)
		{"Hello World!a", "SGVsbG8gV29ybGQhYQ=="},  // multiple of 3 byte length + 1
		{"Hello World!ab", "SGVsbG8gV29ybGQhYWI="}, // multiple of 3 byte length + 2
		{"", ""},      // empty byte slice
		{"6", "Ng=="}, // short text
		{"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.", "TG9yZW0gaXBzdW0gZG9sb3Igc2l0IGFtZXQsIGNvbnNlY3RldHVyIGFkaXBpc2NpbmcgZWxpdCwgc2VkIGRvIGVpdXNtb2QgdGVtcG9yIGluY2lkaWR1bnQgdXQgbGFib3JlIGV0IGRvbG9yZSBtYWduYSBhbGlxdWEuIFV0IGVuaW0gYWQgbWluaW0gdmVuaWFtLCBxdWlzIG5vc3RydWQgZXhlcmNpdGF0aW9uIHVsbGFtY28gbGFib3JpcyBuaXNpIHV0IGFsaXF1aXAgZXggZWEgY29tbW9kbyBjb25zZXF1YXQuIER1aXMgYXV0ZSBpcnVyZSBkb2xvciBpbiByZXByZWhlbmRlcml0IGluIHZvbHVwdGF0ZSB2ZWxpdCBlc3NlIGNpbGx1bSBkb2xvcmUgZXUgZnVnaWF0IG51bGxhIHBhcmlhdHVyLiBFeGNlcHRldXIgc2ludCBvY2NhZWNhdCBjdXBpZGF0YXQgbm9uIHByb2lkZW50LCBzdW50IGluIGN1bHBhIHF1aSBvZmZpY2lhIGRlc2VydW50IG1vbGxpdCBhbmltIGlkIGVzdCBsYWJvcnVtLg=="}, // Long text
	}

	for _, tc := range testCases {
		result := string(Base64Decode(tc.in))
		if result != tc.expected {
			t.Fatalf("Base64Decode(%s) = %s, want %s", tc.in, result, tc.expected)
		}
	}
}

func BenchmarkBase64Decode(b *testing.B) {
	benchmarks := []struct {
		name     string
		expected string
		in       string
	}{
		{"Hello World!", "Hello World!", "SGVsbG8gV29ybGQh"},         // multiple of 3 byte length (multiple of 24-bits)
		{"Hello World!a", "Hello World!a", "SGVsbG8gV29ybGQhYQ=="},   // multiple of 3 byte length + 1
		{"Hello World!ab", "Hello World!ab", "SGVsbG8gV29ybGQhYWI="}, // multiple of 3 byte length + 2
		{"Empty", "", ""},  // empty byte slice
		{"6", "6", "Ng=="}, // short text
		{"Lorem ipsum", "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.", "TG9yZW0gaXBzdW0gZG9sb3Igc2l0IGFtZXQsIGNvbnNlY3RldHVyIGFkaXBpc2NpbmcgZWxpdCwgc2VkIGRvIGVpdXNtb2QgdGVtcG9yIGluY2lkaWR1bnQgdXQgbGFib3JlIGV0IGRvbG9yZSBtYWduYSBhbGlxdWEuIFV0IGVuaW0gYWQgbWluaW0gdmVuaWFtLCBxdWlzIG5vc3RydWQgZXhlcmNpdGF0aW9uIHVsbGFtY28gbGFib3JpcyBuaXNpIHV0IGFsaXF1aXAgZXggZWEgY29tbW9kbyBjb25zZXF1YXQuIER1aXMgYXV0ZSBpcnVyZSBkb2xvciBpbiByZXByZWhlbmRlcml0IGluIHZvbHVwdGF0ZSB2ZWxpdCBlc3NlIGNpbGx1bSBkb2xvcmUgZXUgZnVnaWF0IG51bGxhIHBhcmlhdHVyLiBFeGNlcHRldXIgc2ludCBvY2NhZWNhdCBjdXBpZGF0YXQgbm9uIHByb2lkZW50LCBzdW50IGluIGN1bHBhIHF1aSBvZmZpY2lhIGRlc2VydW50IG1vbGxpdCBhbmltIGlkIGVzdCBsYWJvcnVtLg=="}, // Long text
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Base64Decode(bm.in)
			}
		})
	}
}

func TestBase64EncodeDecodeInverse(t *testing.T) {
	testCases := []struct {
		in string
	}{
		{"Hello World!"},   // multiple of 3 byte length (multiple of 24-bits)
		{"Hello World!a"},  // multiple of 3 byte length + 1
		{"Hello World!ab"}, // multiple of 3 byte length + 2
		{""},               // empty byte slice
		{"6"},              // short text
		{"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."}, // Long text
	}

	for _, tc := range testCases {
		result := string(Base64Decode(Base64Encode([]byte(tc.in))))
		if result != tc.in {
			t.Fatalf("Base64Decode(Base64Encode(%s)) = %s, want %s", tc.in, result, tc.in)
		}
	}
}

func FuzzBase64Encode(f *testing.F) {
	f.Add([]byte("hello"))
	f.Fuzz(func(t *testing.T, input []byte) {
		result := Base64Decode(Base64Encode(input))
		for i := 0; i < len(input); i++ {
			if result[i] != input[i] {
				t.Fatalf("with input '%s' - expected '%s', got '%s' (mismatch at position %d)", input, input, result, i)
			}
		}
	})
}

func TestBase64DecodeMalformedLength(t *testing.T) {
	// A base64 payload is always a multiple of 4 characters. An input of any
	// other length must not index past the end of the string and panic.
	testCases := []struct {
		in       string
		expected string
	}{
		{"SGVsbG8h", "Hello!"}, // well-formed control case
		{"ABCDE", "\x00\x10\x83"},
		{"A", ""},
		{"AB", ""},
		{"ABC", ""},
	}

	for _, tc := range testCases {
		result := string(Base64Decode(tc.in))
		if result != tc.expected {
			t.Fatalf("Base64Decode(%q) = %q, want %q", tc.in, result, tc.expected)
		}
	}
}

func TestBase64EncodeDoesNotMutateInput(t *testing.T) {
	// Encoding a slice whose capacity exceeds its length must not write padding
	// bytes into the caller's backing array.
	backing := []byte{1, 2, 3, 4, 5, 6}
	arg := backing[:4] // len 4, cap 6; 4 % 3 == 1 triggers padding

	Base64Encode(arg)

	want := []byte{1, 2, 3, 4, 5, 6}
	for i := range want {
		if backing[i] != want[i] {
			t.Fatalf("Base64Encode mutated caller backing array: got %v, want %v", backing, want)
		}
	}
}
