package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	w, _ := os.Create("DB.csv")
	wr := csv.NewWriter(bufio.NewWriter(w))
	wr.Write([]string{"A", "1"})
	wr.Write([]string{"B", "2"})
	wr.Flush()
	w.Close()

	r, _ := os.Open("DB.csv")
	read := csv.NewReader(r)
	read.FieldsPerRecord = -1
	data, _ := read.ReadAll()
	for _, e := range data {
		fmt.Println(e[0] + " " + e[1])
	}
	r.Close()
}
