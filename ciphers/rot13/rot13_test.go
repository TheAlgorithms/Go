package main

import (
	"reflect"
	"testing"
)

var testData = []struct{
	description string
	input       string
}{
	{
		"Encrypt letter 'a'",
		"a",
	},
	{
		"Encrypt 'hello world'",
		"hello world",
	},
	{
		"Encrypt full sentence",
		"the quick brown fox jumps over the lazy dog.",
	},
	{
		"Encrypt full sentence from RSAcipher.go main function",
		"i think rsa is really great",
	},
}

func TestRot13(t *testing.T) {
	for _, test := range testData {
		t.Run(test.description, func(t *testing.T) {
			encrypted := rot13(test.input)
			decrypted := rot13(encrypted)
			if actual := decrypted; !reflect.DeepEqual(actual, test.input){
				t.Logf("FAIL: %s", test.description)
				t.Fatalf("Expecting %v, actual %v", test.input, actual)
			}
		})
	}
}
