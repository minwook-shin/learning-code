package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

func main() {
	r, _ := zip.OpenReader("./test=.zip")
	defer r.Close()

	for _, f := range r.File {
		fmt.Printf("Contents of %s:\n", f.Name)
		rc, _ := f.Open()
		io.CopyN(os.Stdout, rc, 10000)
		defer rc.Close()
	}
}
