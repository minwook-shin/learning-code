package main

import "github.com/jung-kurt/gofpdf"

func main() {
	pdf := gofpdf.New("portrait", "mm", "A4", "")

	pdf.AddPage()

	pdf.SetAuthor("author", true)

	pdf.SetTitle("title", true)

	pdf.SetFont("Arial", "B", 16)

	pdf.Cell(50, 20, "Hello, world!")

	err := pdf.OutputFileAndClose("test.pdf")
	if err != nil {
		panic(err)
	}
}
