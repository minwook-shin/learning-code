package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	rInt := '한'
	byteBuffer := make([]byte, 3)

	n := utf8.EncodeRune(byteBuffer, rInt)

	fmt.Println(byteBuffer)
	fmt.Println(n)

	r, size := utf8.DecodeRune(byteBuffer)
	fmt.Printf("%c %v\n", r, size)

	fmt.Println(utf8.FullRune(byteBuffer))
	fmt.Println(utf8.FullRune(byteBuffer[:2]))

	fmt.Println("bytes =", len(byteBuffer))
	fmt.Println("runes =", utf8.RuneCount(byteBuffer))

	fmt.Println(utf8.RuneLen('a'))
	fmt.Println(utf8.RuneLen('한'))

	valid := []byte("Hello, 고 언어")
	fmt.Println(utf8.Valid(valid))
	validRune := 'a'
	fmt.Println(utf8.ValidRune(validRune))
	validString := "Hello, 世界"
	fmt.Println(utf8.ValidString(validString))
}
