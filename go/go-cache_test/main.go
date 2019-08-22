package main

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
)

func main() {
	c := cache.New(1*time.Minute, 2*time.Minute)

	c.Set("test_key_1", "v1", cache.DefaultExpiration)

	c.Set("test_key_2", 1, cache.NoExpiration)
	c.Delete("test_key_2")

	c.SetDefault("test_key_3", 2)
	c.Delete("test_key_3")

	getInterface, foundBool := c.Get("test_key_1")
	if foundBool {
		fmt.Println(getInterface)
	}

	getInterface, foundBool = c.Get("test_key_1")
	if foundBool {
		fmt.Println(getInterface.(string))
	}

	if getInterface, foundBool = c.Get("test_key_1"); foundBool {
		value := getInterface.(string)
		fmt.Println(value)
	}

	var strVar string
	if getInterface, foundBool = c.Get("test_key_1"); foundBool {
		strVar = getInterface.(string)
		fmt.Println(strVar)
	}

	c.SaveFile("test")

	c.Delete("test_key_1")

	c.LoadFile("test")

	if getInterface, foundBool = c.Get("test_key_1"); foundBool {
		strVar = getInterface.(string)
		fmt.Println(strVar)
	}

	c.Flush()
}
