package main

import (
	"archive/zip"
	"path/filepath"
	"io"
	"os"
)

func main() {
	file, _ := os.Create("./test.zip")
	defer file.Close()

	archive := zip.NewWriter(file)
	defer archive.Close()

	filepath.Walk("./db.csv", 
	func(path string, name os.FileInfo, _ error) error {
		header, _ := zip.FileInfoHeader(name)
		header.Method = zip.Deflate
		writer, _ := archive.CreateHeader(header)

		file, error := os.Open(path)
		defer file.Close()
		io.Copy(writer, file)
		return error
	})

}
