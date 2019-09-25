package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	file := excelize.NewFile()
	sheet := file.NewSheet("Sheet1")
	file.SetActiveSheet(sheet)
	file.SetCellValue("Sheet1", "A1", "VALUE")
	err := file.SaveAs("TEST_1.xlsx")
	if err != nil {
		fmt.Println(err)
	}

	items := map[string]string{"A2": "Fisrt", "A3": "Second", "B1": "Java", "C1": "Python"}
	values := map[string]int{"B2": 2, "C2": 4, "B3": 3, "C3": 6}
	sheet = file.NewSheet("Sheet2")
	file.SetActiveSheet(sheet)
	for key, value := range items {
		file.SetCellValue("Sheet2", key, value)
	}
	for key, value := range values {
		file.SetCellValue("Sheet2", key, value)
	}
	err = file.AddChart("Sheet2", "D1", `{"type":"col3DClustered","series":[
		{"name":"Sheet2!$A$2","categories":"Sheet2!$B$1:$C$1","values":"Sheet2!$B$2:$C$2"},
		{"name":"Sheet2!$A$3","categories":"Sheet2!$B$1:$C$1","values":"Sheet2!$B$3:$C$3"}],
		"title":{"name":"Test Chart"}}`)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = file.SaveAs("TEST_1.xlsx")
	if err != nil {
		fmt.Println(err)
	}

	file, err = excelize.OpenFile("TEST_1.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	value, err := file.GetCellValue("Sheet1", "A1")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(value)

	rows, err := file.GetRows("Sheet1")
	for _, row := range rows {
		for _, value := range row {
			fmt.Println(value)
		}
	}
}
