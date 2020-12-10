package main

import (
	"github.com/reaperhero/gitwebhook-go/model"
	"github.com/sirupsen/logrus"
)


func init()  {
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	client := model.NewClientGithub()
	client.SortSearchRepositoryByTopic()
}
