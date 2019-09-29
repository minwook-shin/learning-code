package main

import (
	"compress/flate"
	"fmt"

	"github.com/mholt/archiver"
)

func main() {
	zip := archiver.Zip{
		CompressionLevel:       flate.BestCompression,
		MkdirAll:               true,
		SelectiveCompression:   true,
		ContinueOnError:        false,
		OverwriteExisting:      true,
		ImplicitTopLevelFolder: false,
	}

	files := []string{
		"test_src.md",
	}

	err := zip.Archive(files, "test.zip")
	if err != nil {
		panic(err)
	}

	err = archiver.Extract("test.zip", "test_src.md", "Extract_output")
	if err != nil {
		panic(err)
	}

	err = archiver.Unarchive("test.zip", "Unarchive_output")
	if err != nil {
		panic(err)
	}

	err = archiver.Walk("test.zip", func(f archiver.File) error {
		fmt.Println(f.Name(), f.Size())
		return nil
	})
	if err != nil {
		panic(err)
	}
}
