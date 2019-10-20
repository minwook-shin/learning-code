package t

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestFunc(t *testing.T) {

	convey.Convey("starting value", t, func() {
		x := 0

		convey.Convey("incremented", func() {
			x++

			convey.Convey("ShouldEqual", func() {
				convey.So(x, convey.ShouldEqual, 1)
			})
		})
	})

	convey.Convey("Numeric comparison", t, func() {
		x := 1
		convey.So(x, convey.ShouldBeGreaterThan, 0)
		convey.So(x, convey.ShouldBeLessThan, 2)

		convey.So(x, convey.ShouldBeBetween, 0, 2)
		convey.So(x, convey.ShouldNotBeBetween, 5, 10)
	})

	convey.Convey("Container", t, func() {
		intArr := []int{1, 2, 3}
		intMap := map[int]int{1: 1, 2: 2, 3: 3}

		convey.So(intArr, convey.ShouldContain, 2)
		convey.So(intArr, convey.ShouldNotContain, 4)
		convey.So(intMap, convey.ShouldContainKey, 2)
		convey.So(intMap, convey.ShouldNotContainKey, 4)

		convey.So(1, convey.ShouldBeIn, []int{1, 2, 3})
		convey.So(0, convey.ShouldNotBeIn, []int{1, 2, 3})

		convey.So([]int{}, convey.ShouldBeEmpty)
		convey.So([]int{1}, convey.ShouldNotBeEmpty)
		convey.So([]int{1, 2}, convey.ShouldHaveLength, 2)
	})

	convey.Convey("String", t, func() {
		convey.So("helloworld", convey.ShouldStartWith, "h")
		convey.So("helloworld", convey.ShouldNotStartWith, "e")
		convey.So("helloworld", convey.ShouldEndWith, "world")
		convey.So("helloworld", convey.ShouldNotEndWith, "hello")

		convey.So("helloworld", convey.ShouldContainSubstring, "ow")
		convey.So("helloworld", convey.ShouldNotContainSubstring, "worldhello")
	})

	convey.Convey("Typechecking", t, func() {
		convey.So(10, convey.ShouldHaveSameTypeAs, 0)
		convey.So(10, convey.ShouldNotHaveSameTypeAs, "10")
	})

}
