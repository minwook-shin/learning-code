package main

import (
	"fmt"

	"github.com/robertkrimen/otto"
)

func main() {
	jsRuntime := otto.New()
	jsRuntime.Run(`test = 1 + 1;`)

	if value, err := jsRuntime.Get("test"); err == nil {
		if intValue, err := value.ToInteger(); err == nil {
			fmt.Println(intValue)
		}
	}

	jsRuntime.Set("val", 10)
	jsRuntime.Run(`console.log(val);`)

	jsRuntime.Set("hello", "world")
	jsRuntime.Run(`console.log(hello);`)

	value, err := jsRuntime.Run("hello.length")
	{
		returnValue, _ := value.ToInteger()
		fmt.Println(returnValue)
	}

	value, err = jsRuntime.Run("fake.length")
	if err != nil {
		fmt.Println(err)
	}

	jsRuntime.Set("jsValue", func(call otto.FunctionCall) otto.Value {
		return call.Argument(0)
	})

	result, _ := jsRuntime.Run(`jsValue(1.0);`)
	fmt.Println(result)
}
