package search

import "errors"

// ErrNotFound is returned by search functions when target is not found
var ErrNotFound = errors.New("target not found in array")
