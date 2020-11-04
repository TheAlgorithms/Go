package main

import (
	"reflect"
	"testing"
)

var xorTestData = []struct {
	description string
	input       string
	key         int
	encrypted   string
}{
	{
		"Encrypt letter 'a' with key 0 makes no changes",
		"a",
		0,
		"a",
	},
	{
		"Encrypt letter 'a' with key 1",
		"a",
		1,
		"`",
	},
	{
		"Encrypt letter 'a' with key 10",
		"a",
		10,
		"k",
	},
	{
		"Encrypt 'hello world' with key 0 makes no changes",
		"hello world",
		0,
		"hello world",
	},
	{
		"Encrypt 'hello world' with key 1",
		"hello world",
		1,
		"idmmn!vnsme",
	},
	{
		"Encrypt 'hello world' with key 10",
		"hello world",
		10,
		"boffe*}exfn",
	},
	{
		"Encrypt full sentence with key 64",
		"the quick brown fox jumps over the lazy dog.",
		64,
		"4(%`15)#+`\"2/7.`&/8`*5-03`/6%2`4(%`,!:9`$/'n",
	},
	{
		"Encrypt a word with key 32 make the case swap",
		"abcdefghijklmNOPQRSTUVWXYZ",
		32,
		"ABCDEFGHIJKLMnopqrstuvwxyz",
	},
}

func TestXorCipherEncrypt(t *testing.T) {
	for _, test := range xorTestData {
		t.Run(test.description, func(t *testing.T) {

			message := toASCII([]rune(test.input))
			encrypted := encrypt(test.key, message)

			if actual := decodeToString(encrypted); !reflect.DeepEqual(actual, test.encrypted) {
				t.Logf("FAIL: %s", test.description)
				t.Fatalf("Expecting %s, actual %s", test.encrypted, actual)
			}
		})
	}
}

func TestXorCipherDecrypt(t *testing.T) {
	for _, test := range xorTestData {
		t.Run(test.description, func(t *testing.T) {

			message := toASCII([]rune(test.encrypted))
			decrypted := decrypt(test.key, message)

			if actual := decodeToString(decrypted); !reflect.DeepEqual(actual, test.input) {
				t.Logf("FAIL: %s", test.description)
				t.Fatalf("Expecting %s, actual %s", test.input, actual)
			}
		})
	}
}
