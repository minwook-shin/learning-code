package main

import (
	"fmt"
	"os"
	"os/exec"

	prompt "github.com/c-bata/go-prompt"
)

func executor(text string) {
	fmt.Println(text)

	if text == "bash" {
		cmd := exec.Command("bash")
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	}
	return
}

func completer(text prompt.Document) []prompt.Suggest {
	suggest := []prompt.Suggest{
		{Text: "english", Description: "english Description"},
		{Text: "한국어", Description: "한국어 설명"},

		{Text: "bash"},
	}
	return prompt.FilterHasPrefix(suggest, text.GetWordBeforeCursor(), true)
}

func main() {
	mainPrompt := prompt.New(
		executor,
		completer,
		prompt.OptionTitle("prompt"),
		prompt.OptionHistory([]string{"cd ./"}),
		prompt.OptionPrefixTextColor(prompt.White),
		prompt.OptionPreviewSuggestionTextColor(prompt.Blue),
		prompt.OptionSelectedSuggestionBGColor(prompt.LightGray),
		prompt.OptionSuggestionBGColor(prompt.DarkGray),
	)
	mainPrompt.Run()
}
