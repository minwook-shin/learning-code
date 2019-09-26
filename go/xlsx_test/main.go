package main

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

func main() {
	file := xlsx.NewFile()

	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		panic(err)
	}

	row := sheet.AddRow()
	cell := row.AddCell()
	cell.Value = "VALUE"

	err = file.Save("test.xlsx")
	if err != nil {
		panic(err)
	}

	file, err = xlsx.OpenFile("test.xlsx")
	if err != nil {
		panic(err)
	}
	for _, sheet := range file.Sheets {
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				fmt.Println(cell.String())
			}
		}
	}
}
