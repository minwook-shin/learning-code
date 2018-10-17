package main

import (
	"archive/tar"
	"bytes"
	"io"
	"os"
)

func main() {
	var buffer bytes.Buffer
	tarWriter := tar.NewWriter(&buffer)
	defer tarWriter.Close()

	var files = []struct {
		Name, Body string
	}{
		{"test.txt", "test file"},
	}

	for _, file := range files {
		tarHeader := &tar.Header{
			Name: file.Name,
			Mode: 0600,
			Size: int64(len(file.Body)),
		}
		tarWriter.WriteHeader(tarHeader)
		tarWriter.Write([]byte(file.Body))
	}

	tarReader := tar.NewReader(&buffer)
	for {
		io.Copy(os.Stdout, tarReader)
		_, err := tarReader.Next()
		if err == io.EOF {
			break
		}
	}
}
