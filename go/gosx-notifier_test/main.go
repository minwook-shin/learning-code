package main

import (
	gosxnotifier "github.com/deckarep/gosx-notifier"
)

func main() {
	note := gosxnotifier.NewNotification("메시지")

	note.Title = "제목"
	note.Subtitle = "부제목"

	note.Sound = gosxnotifier.Default

	note.Group = "com.example.app.identifier"

	//note.Sender = "com.apple.Safari"
	note.Link = "https://minwook-shin.github.io"

	note.AppIcon = "original.png"
	note.ContentImage = "original.png"

	err := note.Push()

	if err != nil {
		panic(err)
	}
}
