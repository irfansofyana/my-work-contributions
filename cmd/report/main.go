package main

import (
	"fmt"
	"log"

	"github.com/irfansofyana/workcontributions/pkg/github"
	"github.com/siuyin/dflt"
)

func main() {
	accessToken := dflt.EnvString("GITHUB_TOKEN", "")
	client := github.CreateClient(accessToken)
	issues, err := github.SearchIssues(client, github.SearchIssuesParam{
		GithubUsername:   "irfansofyana",
		GithubOrg:        "xendit",
		CreatedStartFrom: "2022-04-01",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Number of issues found: %d\n", len(issues))
	for _, issue := range issues {
		fmt.Println(*issue.ID, *issue.Title, *issue.URL)
	}
}
