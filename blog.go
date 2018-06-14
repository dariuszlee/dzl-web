package main

import (
	"encoding/json"
	// "fmt"
	"io/ioutil"
	// "bufio"
	// "os"
)

type blog struct {
	Title       string
	Description string
	Image       string
	Id          int
}

func getBlogs(amount int) []blog {
	b1 := []blog{}
	for i := 0; i < amount; i++ {
		b := blog{"First Post", "This is a longer description, etc, etc", "urlImage", i}
		b1 = append(b1, b)
	}
	return b1
}

func readBlogDb() []blog {
	var blogs []blog
	data, _ := ioutil.ReadFile("./markdown/blog.json")
	json.Unmarshal(data, &blogs)
	return blogs
}

func readBlogMarkdown() []string {
	var retData []string
	var str []byte
	line := 0
	f, _ := ioutil.ReadFile("./markdown/welcome.md")
	for i := 0; i < len(f); i++ {
		if f[i] == 10 {
			if len(str) != 0 {
				retData = append(retData, string(str))
				line++
				if line == 3 {
					break
				}
			}
			str = []byte{}
		} else {
			str = append(str, f[i])
		}
	}
	return retData
}

// func main() {
// 	strings := readBlogMarkdown()
// 	fmt.Print(strings)
// }
