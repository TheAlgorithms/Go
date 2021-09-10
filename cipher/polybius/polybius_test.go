package polybius

import (
	"fmt"
	"log"
	"testing"
)

func ExampleNewPolybius() {
	// initialize
	const (
		plainText  = "HogeFugaPiyoSpam"
		size       = 5
		characters = "HogeF"
		key        = "abcdefghijklmnopqrstuvwxy"
	)
	p, err := NewPolybius(key, size, characters)
	if err != nil {
		log.Fatalf("failed NewPolybius: %v", err)
	}
	encryptedText, err := p.Encrypt(plainText)
	if err != nil {
		log.Fatalf("failed Encrypt: %v", err)
	}
	fmt.Printf("Encrypt=> plainText: %s, encryptedText: %s\n", plainText, encryptedText)

	decryptedText, err := p.Decrypt(encryptedText)
	if err != nil {
		log.Fatalf("failed Decrypt: %v", err)
	}
	fmt.Printf("Decrypt=> encryptedText: %s, decryptedText: %s\n", encryptedText, decryptedText)

	// Output:
	// Encrypt=> plainText: HogeFugaPiyoSpam, encryptedText: OGGFOOHFOHFHOOHHEHOEFFGFEEEHHHGG
	// Decrypt=> encryptedText: OGGFOOHFOHFHOOHHEHOEFFGFEEEHHHGG, decryptedText: HOGEFUGAPIYOSPAM
}

func TestNewPolybius(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name       string
		size       int
		characters string
		key        string
		wantErr    string
	}{
		{
			name: "correct initialization", size: 5, characters: "HogeF", key: "abcdefghijklmnopqrstuvwxy", wantErr: "",
		},
		{
			name: "truncate characters", size: 5, characters: "HogeFuga", key: "abcdefghijklmnopqrstuvwxy", wantErr: "",
		},
		{
			name: "invaid key", size: 5, characters: "HogeFuga", key: "abcdefghi", wantErr: "len(key): 9 must be as long as size squared: 25",
		},
		{
			name: "invalid characters", size: 5, characters: "HogeH", key: "abcdefghijklmnopqrstuvwxy", wantErr: "\"chars\" contains same character: H",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewPolybius(tc.key, tc.size, tc.characters)
			if err != nil && err.Error() != tc.wantErr {
				t.Errorf("failed NewPolybius: %v", err)
			}
		})
	}
}

func TestPolybiusEncrypt(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name string
		text string
		want string
	}{
		{
			name: "correct encryption", text: "HogeFugaPiyoSpam", want: "OGGFOOHFOHFHOOHHEHOEFFGFEEEHHHGG",
		},
		{
			name: "invalid encryption", text: "hogz", want: "failed encipher: Z does not exist in keys",
		},
	}
	// initialize
	const (
		size       = 5
		characters = "HogeF"
		key        = "abcdefghijklmnopqrstuvwxy"
	)
	p, err := NewPolybius(key, size, characters)
	if err != nil {
		t.Fatalf("failed NewPolybius: %v", err)
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			encrypted, err := p.Encrypt(tc.text)
			if err != nil {
				if err.Error() != tc.want {
					t.Errorf("failed Encrypt: %v", err)
				}
			} else if encrypted != tc.want {
				t.Errorf("Encrypt: %v, want: %v", encrypted, tc.want)
			}
		})
	}
}

func TestPolybiusDecrypt(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name string
		text string
		want string
	}{
		{
			name: "correct decryption", text: "OGGFOOHFOHFHOOHHEHOEFFGFEEEHHHGG", want: "HOGEFUGAPIYOSPAM",
		},
		{
			name: "invalid decryption(position of even number)", text: "hogz", want: "failed decipher: Z does not exist in characters",
		},
		{
			name: "invalid decryption(position of odd number)", text: "hode", want: "failed decipher: D does not exist in characters",
		},
		{
			name: "invalid text size which is odd", text: "hog", want: "failed decipher: the size of \"chars\" must be even",
		},
	}
	// initialize
	const (
		size       = 5
		characters = "HogeF"
		key        = "abcdefghijklmnopqrstuvwxy"
	)
	p, err := NewPolybius(key, size, characters)
	if err != nil {
		t.Fatalf("failed NewPolybius: %v", err)
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			encrypted, err := p.Decrypt(tc.text)
			if err != nil {
				if err.Error() != tc.want {
					t.Errorf("failed Encrypt: %v", err)
				}
			} else if encrypted != tc.want {
				t.Errorf("Encrypt: %v, want: %v", encrypted, tc.want)
			}
		})
	}
}
