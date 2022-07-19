package main

import (
	"github.com/irfansofyana/workcontributions/internal/domain/infrastructure"
	"github.com/irfansofyana/workcontributions/internal/domain/repository"
	"github.com/irfansofyana/workcontributions/internal/report"
	"github.com/irfansofyana/workcontributions/pkg/github"
	"github.com/siuyin/dflt"
)

func main() {
	accessToken := dflt.EnvString("GITHUB_TOKEN", "")
	githubUsername := dflt.EnvString("GITHUB_USERNAME", "irfansofyana")
	githubOrg := dflt.EnvString("GITHUB_ORG", "xendit")

	client := github.CreateClient(accessToken)
	githubRepo := infrastructure.NewGithubRepository(client, githubUsername, githubOrg)

	report.Print(githubRepo, &repository.WorkRepositoryParam{WorkAfterDate: "2022-04-01"})
}
