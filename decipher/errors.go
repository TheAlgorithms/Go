package decipher

import "errors"

var NoTextToDecryptError = errors.New("no text to cipher")

var KeyMissingError = errors.New("no key provided")
