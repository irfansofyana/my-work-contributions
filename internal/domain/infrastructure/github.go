package infrastructure

import (
	"github.com/google/go-github/github"
	"github.com/irfansofyana/workcontributions/internal/domain/model"
	repo "github.com/irfansofyana/workcontributions/internal/domain/repository"

	gh "github.com/irfansofyana/workcontributions/pkg/github"
)

// GithubRepository struct
type GithubRepository struct {
	GithubClient   *github.Client
	GithubUsername string
	GithubOrg      string
}

// NewGithubRepository is function to create a new GithubRepository instance
func NewGithubRepository(client *github.Client, githubUsername, githubOrg string) GithubRepository {
	return GithubRepository{client, githubUsername, githubOrg}
}

// GetListOfWork function is to get list of work based on Github data
func (gr GithubRepository) GetListOfWork(param *repo.WorkRepositoryParam) ([]model.Work, error) {
	searchIssueParam := &gh.SearchIssuesParam{
		GithubUsername: gr.GithubUsername,
		GithubOrg:      gr.GithubOrg,
	}
	if param != nil {
		searchIssueParam.CreatedStartFrom = param.WorkAfterDate
	}

	githubIssues, err := gh.SearchIssues(gr.GithubClient, searchIssueParam)
	if err != nil {
		return []model.Work{}, err
	}

	return listGithubIssuesToListOfWork(githubIssues), nil
}

func listGithubIssuesToListOfWork(issues []github.Issue) []model.Work {
	listOfWork := make([]model.Work, 0)

	for _, issue := range issues {
		listOfWork = append(listOfWork, githubIssueToWork(issue))
	}

	return listOfWork
}

func githubIssueToWork(issue github.Issue) model.Work {
	work := model.Work{
		What:        *issue.Title,
		RelatedLink: *issue.URL,
		Type:        model.GithubIssue,
	}

	if issue.IsPullRequest() {
		work.Type = model.GithubPR
	}

	return work
}
