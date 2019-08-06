package main

import (
	"fmt"
	"time"

	"github.com/gosuri/uilive"
)

func main() {
	writer := uilive.New()
	writer.Start()

	for i := 0; i <= 100; i++ {
		fmt.Fprintf(writer, "Downloading.. %d\n", i)
		fmt.Fprintf(writer.Newline(), "Downloading.. %d\n", i)

		time.Sleep(time.Millisecond * 5)
	}

	fmt.Fprintln(writer, "Finished")

	fmt.Fprintf(writer.Bypass(), "Downloaded files\n")

	writer.Stop()
}
