package cipher

import "errors"

// ErrorFailedToEncrypt Raised when a cipher function fails to encrypt the message
var ErrorFailedToEncrypt = errors.New("failed to Encrypt")

var NoTextToEncryptError = errors.New("no text to cipher")

var KeyMissingError = errors.New("no key provided")
