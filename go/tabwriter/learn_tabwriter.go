package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func main() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 0, '*', tabwriter.Debug)
	fmt.Fprintln(w, "a\tb\tc")
	fmt.Fprintln(w, "aa\tbb\tcc")
	fmt.Fprintln(w, "aaa\t\t")
	fmt.Fprintln(w, "aaaa\tbbbb\tcccc")
	w.Flush()

	w = tabwriter.NewWriter(os.Stdout, 0, 0, 5, '*', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintln(w, "a\tb\tc")
	fmt.Fprintln(w, "aa\tbb\tcc")
	fmt.Fprintln(w, "aaa\t\t")
	fmt.Fprintln(w, "aaaa\tbbbb\tcccc")
	w.Flush()
}
