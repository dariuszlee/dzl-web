package main

import (
	"fmt"
	"gopkg.in/russross/blackfriday.v2"
	"io/ioutil"
)

func readBlog(fileName string) string {
	f, _ := ioutil.ReadFile(fileName)
	output := string(blackfriday.Run(f))
	return output
}

func formatBlog(fileName string) string {
	blogRaw := readBlog(fileName)
	return fmt.Sprintf("{{define \"blog\"}}%s{{end}}", blogRaw)
}
