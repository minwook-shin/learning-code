package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		if scanner.Text() == "!q" {
			os.Exit(0)
		}

		fmt.Println(strings.Contains(scanner.Text(), "hello"))
		fmt.Println(strings.Count(scanner.Text(), "l"))
		fmt.Println(strings.Index(scanner.Text(), "hello"))
		fmt.Println(strings.Repeat(scanner.Text(), 3))
		fmt.Println(strings.Split(scanner.Text(), ","))
		fmt.Println(strings.ToUpper(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		os.Exit(1)
	}
}
