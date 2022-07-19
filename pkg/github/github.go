package github

import (
	"context"
	"fmt"
	"time"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type SearchIssuesParam struct {
	GithubUsername   string
	GithubOrg        string
	CreatedStartFrom string // should be in "yyyy-mm-dd" format
}

// CreateClient is function to create a new instance of github client
func CreateClient(accessToken string) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(context.Background(), ts)

	return github.NewClient(tc)
}

// SearchIssues is function to find issues with the requested param
func SearchIssues(client *github.Client, param SearchIssuesParam) ([]github.Issue, error) {
	query := fmt.Sprintf("author:%s org:xendit", param.GithubUsername)
	if param.CreatedStartFrom != "" {
		query = fmt.Sprintf("%s created:>=%s", query, param.CreatedStartFrom)
	}

	issues := make([]github.Issue, 0)

	opt := &github.SearchOptions{
		ListOptions: github.ListOptions{
			Page:    1,
			PerPage: 100,
		},
	}

	for {
		issuesSearchResult, response, err := client.Search.Issues(context.Background(), query, opt)
		if err != nil {
			return []github.Issue{}, err
		}

		issues = append(issues, issuesSearchResult.Issues...)

		if response.LastPage == 0 {
			break
		}

		opt.ListOptions.Page = response.NextPage

		if response.Remaining == 0 {
			time.Sleep(time.Until(response.Reset.Time)) // to handle too many request
		}
	}

	return issues, nil
}
