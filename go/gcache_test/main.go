package main

import (
	"fmt"
	"time"

	"github.com/bluele/gcache"
)

func main() {
	gc := gcache.New(20).
		LRU().
		Build()
	gc.Set("key", "value")
	value, err := gc.Get("key")
	if err != nil {
		panic(err)
	}
	fmt.Println(value)

	gc.SetWithExpire("key", "value", time.Second*3)
	value, err = gc.Get("key")
	if err != nil {
		panic(err)
	}
	fmt.Println("Get:", value)

	time.Sleep(time.Second * 5)

	value, err = gc.Get("key")
	if err != nil {
		fmt.Println(err)
	}

	gc = gcache.New(20).
		LRU().
		LoaderFunc(func(key interface{}) (interface{}, error) {
			return "load", nil
		}).
		Build()
	value, err = gc.Get("key")
	if err != nil {
		panic(err)
	}
	fmt.Println(value)

	gc = gcache.New(20).
		LRU().
		LoaderExpireFunc(func(key interface{}) (interface{}, *time.Duration, error) {
			expire := 1 * time.Second
			return "ok", &expire, nil
		}).
		EvictedFunc(func(key, value interface{}) {
			fmt.Println("evicted")
		}).
		PurgeVisitorFunc(func(key, value interface{}) {
			fmt.Println("purged")
		}).
		Build()
	value, err = gc.Get("key")
	if err != nil {
		panic(err)
	}
	fmt.Println(value)
	time.Sleep(1 * time.Second)
	value, err = gc.Get("key")
	if err != nil {
		panic(err)
	}
	fmt.Println(value)
	gc.Purge()
}
