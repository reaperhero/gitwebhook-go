package service

import (
	"github.com/google/go-github/v33/github"
)

type GitService interface {
	SortSearchRepositoryByTopic(topicName string) <-chan bool
	ListSummaryOrganization(origin string) (result string, err error)
	GetRepositoryDetail(repo string) (result string, err error) // ("/google/go-github")
}

type clientGithub struct {
	GithubClient *github.Client
}

func NewClientGithub() GitService {
	return &clientGithub{
		GithubClient: github.NewClient(nil),
	}
}
