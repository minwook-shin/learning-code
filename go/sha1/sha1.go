package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	str := "hello,world!"
	hash := sha1.New()
	hash.Write([]byte(str))
	byteString := hash.Sum(nil)

	fmt.Println(str)
	fmt.Printf("%x\n", byteString)
}
