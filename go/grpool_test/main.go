package main

import (
	"fmt"

	"github.com/ivpusic/grpool"
)

func main() {
	pool := grpool.NewPool(100, 50)

	pool.WaitCount(100)

	for i := 0; i < 100; i++ {
		value := i

		pool.JobQueue <- func() {
			defer pool.JobDone()
			fmt.Println(value)
		}
	}
	pool.WaitAll()

	defer pool.Release()
}
