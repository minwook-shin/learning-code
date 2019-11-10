package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/h2non/filetype"
)

func main() {
	bufferByte, _ := ioutil.ReadFile("original.png")

	fileType, _ := filetype.Match(bufferByte)
	if fileType == filetype.Unknown {
		panic(filetype.Unknown)
	}

	fmt.Println(fileType.Extension, fileType.MIME.Value)

	if filetype.IsImage(bufferByte) {
		fmt.Println("image")
	}

	if filetype.IsSupported("png") {
		fmt.Println("supported")
	}

	if filetype.IsMIMESupported("image/png") {
		fmt.Println("MIME supported")
	}

	file, _ := os.Open("original.png")

	header := make([]byte, 8)
	file.Read(header)
	fmt.Println(header)

	if filetype.IsImage(header) {
		fmt.Println("image")
	}
}
