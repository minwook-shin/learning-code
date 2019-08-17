package main

import (
	"fmt"
	"time"

	"github.com/allegro/bigcache"
)

func main() {
	defaultCache, err := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
	if err != nil {
		panic(err)
	}

	defaultCache.Set("test_id", []byte("hello"))

	if entry, err := defaultCache.Get("test_id"); err == nil {
		fmt.Println(string(entry))
	}

	config := bigcache.Config{
		Shards:             1024,
		LifeWindow:         10 * time.Minute,
		CleanWindow:        1 * time.Minute,
		MaxEntriesInWindow: 1000 * 10 * 60,
		MaxEntrySize:       500,
		Verbose:            true,
		HardMaxCacheSize:   0,
		OnRemove:           nil,
		OnRemoveWithReason: nil,
	}

	customCache, err := bigcache.NewBigCache(config)
	if err != nil {
		panic(err)
	}

	customCache.Set("test_id", []byte("hello"))

	if entry, err := customCache.Get("test_id"); err == nil {
		fmt.Println(string(entry))
	}
}
