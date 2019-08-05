package main

import (
	"os"
	"text/tabwriter"

	"github.com/cheynewallace/tabby"
)

func main() {
	table1 := tabby.New()
	table1.AddHeader("HEADER1", "HEADER2", "HEADER3")
	table1.AddLine("first", "second", "third")
	table1.Print()

	table2 := tabby.New()
	table2.AddLine("LINE1:", "content_1", "content_1")
	table2.AddLine("LINE2:", "content_2", "content_2")
	table2.AddLine("LINE3:", "content_3", "content_3")
	table2.Print()

	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, '|', 0)
	table3 := tabby.NewCustom(writer)
	table3.AddLine("LINE1:", "content_1", "content_1")
	table3.AddLine("LINE2:", "content_2", "content_2")
	table3.AddLine("LINE3:", "content_3", "3")
	table3.Print()
}
