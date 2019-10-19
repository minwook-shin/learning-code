package main

import (
	"fmt"
	"math"
	"sort"

	"github.com/google/go-cmp/cmp"
)

type wifi struct {
	SSID string
	IP   string
}

func wifiList() (x, y wifi) {
	x = wifi{
		SSID: "Free_WIFI",
		IP:   "192, 168, 0, 1",
	}
	y = wifi{
		SSID: "Free_WIFI",
		IP:   "192, 168, 0, 2",
	}
	return x, y
}

func main() {
	original, get := wifiList()

	if diff := cmp.Diff(original, get); diff != "" {
		fmt.Println(diff)
	}

	option := cmp.Comparer(func(x, y float64) bool {
		return (math.IsNaN(x) && math.IsNaN(y)) || x == y
	})

	x := []float64{1, math.NaN(), math.E, 0.0}
	y := []float64{1, math.NaN(), math.E, 0.0}

	fmt.Println(cmp.Equal(x, y, option))

	option = cmp.Transformer("Sort", func(i []int) []int {
		arr := append([]int(nil), i...)
		sort.Ints(arr)
		return arr
	})

	originArr := struct{ Ints []int }{[]int{0, 1, 2, 3, 4, 5}}
	copyArr := struct{ Ints []int }{[]int{2, 0, 1, 4, 3, 5}}

	fmt.Println(cmp.Equal(originArr, copyArr, option))
}
