package main

import (
	"bytes"
	"compress/gzip"
	"io"
	"os"
)

func main() {
	var buffer bytes.Buffer
	gzipWriter := gzip.NewWriter(&buffer)

	gzipWriter.Write([]byte("test text"))
	gzipWriter.Close()

	gzipReader, _ := gzip.NewReader(&buffer)
	defer gzipReader.Close()

	io.Copy(os.Stdout, gzipReader)
}
