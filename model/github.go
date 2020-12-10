package model

import (
	"context"
	"github.com/google/go-github/v33/github"
	"github.com/sirupsen/logrus"
)

type clientGithub struct {
	GithubClient *github.Client
}

func NewClientGithub() *clientGithub {
	return &clientGithub{
		GithubClient: github.NewClient(nil),
	}
}

func (g *clientGithub) SearchRepositoryByTopic() {
	topics, _, err := g.GithubClient.Search.Topics(context.Background(), "kubernetes", nil)
	if err != nil {
		logrus.Warn(err)
	}
	logrus.Info(topics.Topics[0].GetDisplayName())
}

func (g *clientGithub) SortSearchRepositoryByTopic() {
	opts := &github.SearchOptions{Sort: "stars", Order: "desc", ListOptions: github.ListOptions{Page: 1, PerPage: 100}}
	result, _, err := g.GithubClient.Search.Repositories(context.Background(), "golang", opts)
	if err != nil {
		logrus.Warn(err)
		return
	}
	markdown := NewGitMarkdown("README.md")
	for _, repository := range result.Repositories {
		markdown.WriteProjectInfo(*repository.Name , *repository.CloneURL , *repository.StargazersCount)
	}
	markdown.File.Close()
}
