package main

import (
	"encoding/json"
	"github.com/reaperhero/gitwebhook-go/service"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	jsonconfig = config{}
	client     = service.NewClientGithub()
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
	count := len(topics)
	wg.Add(count)
	for _, topic := range topics {
		go func() {
			for range client.SortSearchRepositoryByTopic(topic) {
				wg.Done()
				count--
				logrus.Info("wait for " + strconv.Itoa(count*10) + " second")
			}
		}()
		time.Sleep(time.Second * 10)
	}
}

func main() {
	syncTask()
}

func syncTask() {
	finish := sync.WaitGroup{}
	createMarkdown(jsonconfig.Topic, &finish)
	finish.Wait()
}
