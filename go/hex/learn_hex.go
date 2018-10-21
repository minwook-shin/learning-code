package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	source := []byte("Hello world!")

	encodeText := make([]byte, hex.EncodedLen(len(source)))
	hex.Encode(encodeText, source)
	source = make([]byte, len(source))
	fmt.Println(string(encodeText))

	source = make([]byte, hex.DecodedLen(len(encodeText)))
	hex.Decode(source, encodeText)
	fmt.Println(string(source[:]))
}
