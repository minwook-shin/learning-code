package main

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"os"
)

func main() {
	var buffer bytes.Buffer
	writer := zlib.NewWriter(&buffer)
	writer.Write([]byte("test"))
	writer.Close()
	
	fmt.Println(buffer.Bytes())

	reader, _ := zlib.NewReader(&buffer)
	defer reader.Close()
	io.Copy(os.Stdout, reader)
}
