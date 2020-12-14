package main

import (
	"encoding/json"
	"github.com/reaperhero/gitwebhook-go/model"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"sync"
	"time"
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

func createMarkdown(topics []string, wg *sync.WaitGroup) {
	wg.Add(len(topics))
	for _, topic := range topics {
		go func() {
			for range client.SortSearchRepositoryByTopic(topic) {
				wg.Done()
			}
		}()
		time.Sleep(time.Second * 10)
	}
}

func main() {
	finish := sync.WaitGroup{}
	createMarkdown(jsonconfig.Topic, &finish)
	finish.Wait()
}
