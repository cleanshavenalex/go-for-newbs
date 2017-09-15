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
	opt := &github.RepositoryListByOrgOptions{Type: "private"}
	repos, _, err := client.Repositories.ListByOrg(ctx, "dollarshaveclub", opt)
	if err != nil {
		fmt.Printf("Error: %v", err)
	} else {
		fmt.Println("No Error when client.Repositories.List")
		fmt.Printf("dollarshaveclub has %v repos", len(repos))
		// for _, r := range repos {
		// 	fmt.Println(r)
		// }

	}

}
