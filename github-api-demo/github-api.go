package main

import "github.com/google/go-github/github"
import "golang.org/x/oauth2"
import (
	"context"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Starting Github GO API Client")

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	repos, _, err := client.Repositories.ListByOrg(ctx, "dollarshaveclub", nil)
	if err != nil {
		fmt.Printf("Error: %v", err)
	} else {
		fmt.Println("No Error when client.Repositories.List")
		fmt.Println(repos)
		for _, r := range repos {
			fmt.Println(r)
		}

	}

}
