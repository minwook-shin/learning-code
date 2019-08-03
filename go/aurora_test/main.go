package main

import (
	"fmt"

	. "github.com/logrusorgru/aurora"
)

func main() {
	fmt.Println(Bold("Bold"))
	fmt.Println(Faint("Faint"))
	fmt.Println(DoublyUnderline("DoublyUnderline"))
	fmt.Println(Fraktur("Fraktur"))
	fmt.Println(Italic("Italic"))
	fmt.Println(Underline("Underline"))
	fmt.Println(SlowBlink("SlowBlink"))
	fmt.Println(RapidBlink("RapidBlink"))
	fmt.Println(Blink("Blink"))
	fmt.Println(Reverse("Reverse"))
	fmt.Println(Inverse("Inverse"))
	fmt.Println(Conceal("Conceal"))
	fmt.Println(Hidden("Hidden"))
	fmt.Println(CrossedOut("CrossedOut"))
	fmt.Println(StrikeThrough("StrikeThrough"))
	fmt.Println(Framed("Framed"))
	fmt.Println(Encircled("Encircled"))
	fmt.Println(Overlined("Overlined"))

	fmt.Println(Black("Black"))
	fmt.Println(Red("Red"))
	fmt.Println(Green("Green"))
	fmt.Println(Yellow("Yellow"))
	fmt.Println(Brown("Brown"))
	fmt.Println(Blue("Blue"))
	fmt.Println(Magenta("Magenta"))
	fmt.Println(Cyan("Cyan"))
	fmt.Println(White("White"))
	fmt.Println(BrightBlack("BrightBlack"))
	fmt.Println(BrightRed("BrightRed"))
	fmt.Println(BrightGreen("BrightGreen"))
	fmt.Println(BrightYellow("BrightYellow"))
	fmt.Println(BrightBlue("BrightBlue"))
	fmt.Println(BrightMagenta("BrightMagenta"))
	fmt.Println(BrightCyan("BrightCyan"))
	fmt.Println(BrightWhite("BrightWhite"))

	fmt.Println(BgBlack("BgBlack"))
	fmt.Println(BgRed("BgRed"))
	fmt.Println(BgGreen("BgGreen"))
	fmt.Println(BgYellow("BgYellow"))
	fmt.Println(BgBrown("BgBrown"))
	fmt.Println(BgBlue("BgBlue"))
	fmt.Println(BgMagenta("BgMagenta"))
	fmt.Println(BgCyan("BgCyan"))
	fmt.Println(BgWhite("BgWhite"))
	fmt.Println(BgBrightBlack("BgBrightBlack"))
	fmt.Println(BgBrightRed("BgBrightRed"))
	fmt.Println(BgBrightGreen("BgBrightGreen"))
	fmt.Println(BgBrightYellow("BgBrightYellow"))
	fmt.Println(BgBrightBlue("BgBrightBlue"))
	fmt.Println(BgBrightMagenta("BgBrightMagenta"))
	fmt.Println(BgBrightCyan("BgBrightCyan"))
	fmt.Println(BgBrightWhite("BgBrightWhite"))

	hello := Blue("hello world").Bold().BgBrightWhite()
	fmt.Println(hello)

	fmt.Println(
		Gray(0, " 1 ").BgGray(23),
		Gray(3, " 2 ").BgGray(19),
		Gray(7, " 3 ").BgGray(15),
		Gray(11, " 4 ").BgGray(11),
		Gray(15, " 5 ").BgGray(9),
		Gray(19, " 6 ").BgGray(3),
		Gray(23, " 7 ").BgGray(0),
	)
}
