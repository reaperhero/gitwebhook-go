package main

import (
	"encoding/json"
	"github.com/reaperhero/gitwebhook-go/model"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

var (
	jsonconfig = config{}
	client     = model.NewClientGithub()
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	jsonFile, err := os.Open("./config.json")
	if err != nil {
		logrus.Fatal("file not exist..")
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &jsonconfig)
}

func createMarkdown(topics []string) {
	for _, topic := range topics {
		go client.SortSearchRepositoryByTopic(topic)
	}
}
func main() {
	createMarkdown(jsonconfig.Topic)
}
