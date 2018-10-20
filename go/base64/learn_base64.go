package main

import (
	"encoding/base64"
	"os"
	"fmt"
)

func main() {
	str := []byte("hello,world!")
	encoder := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	encoder.Write(str)
	defer encoder.Close()

	base64Str := "aGVsbG8sd29ybGQh"
	str, _ = base64.StdEncoding.DecodeString(base64Str)
	fmt.Printf("%q\n", str)
}
