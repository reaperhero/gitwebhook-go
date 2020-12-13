package model

import (
	"os"
	"strconv"
)

type gitMarkdown struct {
	File     *os.File
	HasTitle bool
}

func NewGitMarkdown(filename string) *gitMarkdown {
	f, _ := os.Create(filename)
	f.WriteString("# topic star desc \n\n\n\n\n")
	return &gitMarkdown{
		File: f,
	}
}

func (m *gitMarkdown) WriteContext(context string) {
	m.File.WriteString(context)
	m.File.Sync()
}

func (m *gitMarkdown) WriteProjectInfo(name, url string, stars int) {
	star := ""
	if stars > 999 {
		star = strconv.Itoa(stars/1000.0) + "k"
	}

	if !m.HasTitle {
		m.File.WriteString("|  仓库   | stars  | \n|-----|-------| \n")
		m.HasTitle = true
	}
	m.File.WriteString("|[" + name + "](" + url + ")" + "|" + star + "|\n")
	m.File.Sync()
}
