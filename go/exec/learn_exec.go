package main

import (
	"fmt"
	"os/exec"
)

func main() {
	path, _ := exec.LookPath("go")
	fmt.Println(path)

	out, _ := exec.Command("date").Output()
	fmt.Println(string(out))

	cmd := exec.Command("sleep", "1")
	cmd.Start()
	cmd.Wait()

	cmd = exec.Command("sh", "-c", "echo hello,world!")
	stdoutStderr, _ := cmd.CombinedOutput()
	fmt.Println(string(stdoutStderr))
}