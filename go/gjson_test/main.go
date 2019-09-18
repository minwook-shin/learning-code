package main

import (
	"fmt"
	"strings"

	"github.com/tidwall/gjson"
)

const json = `{"map":{"first":"hello","last":"world"},
"list": ["first","second","third"],
"int":10}`

func main() {
	gjson.AddModifier("custom", func(json, arg string) string {
		if arg == "upper" {
			return strings.ToUpper(json)
		}
		if arg == "lower" {
			return strings.ToLower(json)
		}
		return json
	})
	result := gjson.Get(json, "map.last")
	fmt.Println(result.String())

	result = gjson.Get(json, "map.last|@custom:upper")
	fmt.Println(result.String())

	gjson.ForEachLine(json, func(line gjson.Result) bool {
		fmt.Println(line.String())
		return true
	})

	result = gjson.Get(json, "list")
	for _, name := range result.Array() {
		fmt.Println(name.String())
	}

	result = gjson.Get(json, "map")
	result.ForEach(func(key, value gjson.Result) bool {
		fmt.Println(value.String())
		return true
	})

	fmt.Println(gjson.Parse(json).Get("map").Get("last"))

	result = gjson.Get(json, "map.last")
	if !result.Exists() {
		panic("no last map")
	}

	if !gjson.Valid(json) {
		panic("invalid json")
	}
	content, ok := gjson.Parse(json).Value().(map[string]interface{})
	if !ok {
		panic("err")
	}
	fmt.Println(content)
}
