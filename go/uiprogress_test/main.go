package main

import (
	"sync"
	"time"

	"github.com/gosuri/uiprogress"
)

func main() {
	uiprogress.Start()
	bar := uiprogress.AddBar(100)

	bar.AppendCompleted()
	bar.PrependElapsed()

	for bar.Incr() {
		time.Sleep(time.Millisecond * 10)
	}

	customBar := uiprogress.AddBar(4)

	customBar.PrependFunc(func(bar *uiprogress.Bar) string {
		return []string{"downloading", "removing", "updating", "running"}[bar.Current()-1]
	})

	for customBar.Incr() {
		time.Sleep(time.Millisecond * 500)
	}

	var wg sync.WaitGroup

	bar1 := uiprogress.AddBar(100).AppendCompleted().PrependElapsed()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for bar1.Incr() {
			time.Sleep(time.Millisecond * 50)
		}
	}()

	bar2 := uiprogress.AddBar(100).PrependElapsed().AppendCompleted()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= bar2.Total; i++ {
			bar2.Set(i)
			time.Sleep(time.Millisecond * 50)
		}
	}()

	wg.Wait()
}
