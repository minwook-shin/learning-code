package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
)

func main() {
	client := github.NewClient(nil)

	organizations, _, err := client.Organizations.List(context.Background(), "minwook-shin", nil)
	if err != nil {
		panic(err)
	}

	for _, organization := range organizations {
		option := &github.RepositoryListByOrgOptions{Type: "public"}
		repositories, _, err := client.Repositories.ListByOrg(context.Background(), organization.GetLogin(), option)
		if err != nil {
			panic(err)
		}

		for index, repository := range repositories {
			fmt.Println(index+1, repository.GetName())
		}
	}
}
