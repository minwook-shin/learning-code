package main

import (
	"fmt"
	"time"

	"github.com/briandowns/spinner"
)

func main() {
	spinnerInstance := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	spinnerInstance.Start()
	time.Sleep(2 * time.Second)
	spinnerInstance.Stop()

	spinnerInstance.Reverse()
	spinnerInstance.Restart()
	time.Sleep(2 * time.Second)
	spinnerInstance.Stop()

	spinnerInstance.UpdateCharSet(spinner.CharSets[36])
	spinnerInstance.Restart()
	time.Sleep(2 * time.Second)
	spinnerInstance.Stop()

	spinnerInstance.UpdateSpeed(1 * time.Millisecond)
	spinnerInstance.Restart()
	time.Sleep(2 * time.Second)
	spinnerInstance.Stop()

	cs := []string{"o", "O"}
	spinnerInstance = spinner.New(cs, 100*time.Millisecond)
	spinnerInstance.Start()
	time.Sleep(2 * time.Second)
	spinnerInstance.Stop()

	spinnerInstance.Prefix = "prefix: "
	spinnerInstance.Suffix = " :suffix"
	spinnerInstance.Start()
	time.Sleep(2 * time.Second)
	spinnerInstance.Stop()

	spinnerInstance.Color("blue", "bold")
	spinnerInstance.Start()
	time.Sleep(2 * time.Second)
	spinnerInstance.Stop()

	cs = spinner.GenerateNumberSequence(10)
	spinnerInstance = spinner.New(cs, 100*time.Millisecond)
	spinnerInstance.Start()
	time.Sleep(2 * time.Second)
	fmt.Println(spinnerInstance.Active())
	spinnerInstance.Stop()

	spinnerInstance = spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	spinnerInstance.FinalMSG = "first line\nsecond line\n"
	spinnerInstance.Start()
	time.Sleep(2 * time.Second)
	spinnerInstance.Stop()
}
