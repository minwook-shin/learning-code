package main

import (
	"fmt"
	"time"

	"github.com/muesli/cache2go"
)

type customStruct struct {
	text  string
	extra []byte
}

func main() {
	cache := cache2go.Cache("cacheTable")

	val := customStruct{"text", []byte{'a'}}
	cache.Add("TestKey", 2*time.Second, &val)

	// cache.Add("TestKey", 0, &val)

	res, err := cache.Value("TestKey")
	if err == nil {
		text := res.Data().(*customStruct).text
		extra := res.Data().(*customStruct).extra
		fmt.Println(text, extra)
	} else {
		panic(err)
	}

	time.Sleep(3 * time.Second)

	res, err = cache.Value("TestKey")
	if err != nil {
		fmt.Println("not cached.")
	}

	cache.Add("TestKey", 0, &val)

	cache.SetAboutToDeleteItemCallback(func(e *cache2go.CacheItem) {
		key := e.Key()
		TextData := e.Data().(*customStruct).text
		ExtraData := e.Data().(*customStruct).extra
		created := e.CreatedOn()
		accessed := e.AccessedOn()
		fmt.Println(key, TextData, ExtraData, created, accessed)
	})

	cache.Delete("TestKey")

	cache.Flush()
}
