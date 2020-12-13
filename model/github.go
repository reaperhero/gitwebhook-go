package model

import (
	"context"
	"github.com/google/go-github/v33/github"
	"github.com/sirupsen/logrus"
)

const (
	exportDir = "report/"
)

type clientGithub struct {
	GithubClient *github.Client
}

func NewClientGithub() *clientGithub {
	return &clientGithub{
		GithubClient: github.NewClient(nil),
	}
}

func (g *clientGithub) SortSearchRepositoryByTopic(topicName string) <-chan bool {
	out := make(chan bool)
	defer close(out)
	opts := &github.SearchOptions{Sort: "stars", Order: "desc", ListOptions: github.ListOptions{Page: 1, PerPage: 100}}
	result, _, err := g.GithubClient.Search.Repositories(context.Background(), topicName, opts)
	if err != nil {
		logrus.Warn(err)
		out <- false
		return out
	}
	markdown := NewGitMarkdown(exportDir + topicName + ".md")
	for _, repository := range result.Repositories {
		markdown.WriteProjectInfo(*repository.Name, *repository.CloneURL, *repository.StargazersCount)
	}
	markdown.File.Close()
	out <- true
	return out
}


func (g *clientGithub) SortSearchRepositoryByFollow(topicName string) <-chan bool {
	out := make(chan bool)
	defer close(out)
	opts := &github.SearchOptions{Sort: "stars", Order: "desc", ListOptions: github.ListOptions{Page: 1, PerPage: 100}}
	result, _, err := g.GithubClient.Search.Repositories(context.Background(), topicName, opts)
	if err != nil {
		logrus.Warn(err)
		out <- false
		return out
	}
	markdown := NewGitMarkdown(exportDir + topicName + ".md")
	for _, repository := range result.Repositories {
		markdown.WriteProjectInfo(*repository.Name, *repository.CloneURL, *repository.StargazersCount)
	}
	markdown.File.Close()
	out <- true
	return out
}
