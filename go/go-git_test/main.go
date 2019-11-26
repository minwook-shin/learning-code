package main

import (
	"fmt"
	"os"

	"gopkg.in/src-d/go-git.v4"
)

func main() {
	_, err := git.PlainClone("test", false, &git.CloneOptions{
		URL:      "https://github.com/github/gitignore",
		Progress: os.Stdout,
	})
	if err != nil {
		fmt.Println(err)
	}

	repository, err := git.PlainOpen("test")
	if err != nil {
		fmt.Println(err)
	}

	worktree, err := repository.Worktree()
	if err != nil {
		panic(err)
	}

	err = worktree.Pull(&git.PullOptions{RemoteName: "origin"})
	if err != nil {
		fmt.Println(err)
	}

	reference, err := repository.Head()
	if err != nil {
		panic(err)
	}
	commit, err := repository.CommitObject(reference.Hash())
	if err != nil {
		panic(err)
	}
	fmt.Println(commit)
}
