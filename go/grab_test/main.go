package main

import (
	"fmt"
	"time"

	"github.com/cavaliercoder/grab"
)

func main() {
	client := grab.NewClient()
	request, _ := grab.NewRequest(".", "http://www.oreilly.com/programming/free/files/functional-programming-python.pdf")

	fmt.Println(request.URL())

	response := client.Do(request)
	fmt.Println(response.HTTPResponse.Status)

	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

Loop:
	for {
		select {
		case <-ticker.C:
			fmt.Printf("%v / %v (%.2f%%)\n",
				response.BytesComplete(),
				response.Size(),
				100*response.Progress())

		case <-response.Done:
			break Loop
		}
	}

	err := response.Err()
	if err != nil {
		panic(err)
	}

	fmt.Println("Saved : ", response.Filename)
}
