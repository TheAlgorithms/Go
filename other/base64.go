package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	text := "Hello, World !"
	enc := base64.StdEncoding.EncodeToString([]byte(text))
	fmt.Println(text)
	fmt.Println(enc)
}
