package service

import (
	"context"
	"errors"
	"github.com/google/go-github/v33/github"
	"github.com/reaperhero/gitwebhook-go/model"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

const (
	exportDir  = "report/"
	Accept     = "application/vnd.github.v3+json"
	ApiAddress = "https://api.github.com"
)

func (g *clientGithub) ListSummaryOrganization(origin string) (result string, err error) {
	url := ApiAddress + "/orgs/" + origin + "/repos"
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return "", errors.New("get error " + err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	return string(body), nil
}

// GetRepositoryDetail("/octokit/octokit.rb")
func (g *clientGithub) GetRepositoryDetail(repo string) (result string, err error) {
	url := ApiAddress + "/repos" + repo
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return "", errors.New("get error " + err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	return string(body), nil
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
	markdown := model.NewGitMarkdown(exportDir + topicName + ".md")
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
	markdown := model.NewGitMarkdown(exportDir + topicName + ".md")
	for _, repository := range result.Repositories {
		markdown.WriteProjectInfo(*repository.Name, *repository.CloneURL, *repository.StargazersCount)
	}
	markdown.File.Close()
	out <- true
	return out
}
