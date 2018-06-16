package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"strings"
)

type blog struct {
	Title       string
	Description string
	Path        string
	Image       string
	Id          int
	Count       int
}

var blogs map[int]blog

// Get latest only
func getBlogs() {
	blogList := readBlogDb()
	blogs = make(map[int]blog)
	for _, blog := range blogList {
		log.Println("Loading blog ", blog)
		blog.Description = getBlogText(150, blog.Path)
		blogs[blog.Id] = blog
	}
}

func getBlogText(numOfBytes int, pathToBlog string) string {
	blog := readBlog(pathToBlog)
	blogReader := strings.NewReader(blog)
	tokens := html.NewTokenizer(blogReader)

	var text bytes.Buffer
	for {
		t := tokens.Next()
		if t == html.ErrorToken {
			break
		} else if t == html.TextToken {
			text.WriteString(tokens.Token().Data)
		} else {
		}

		if text.Len() > numOfBytes {
			text.Truncate(numOfBytes)
			text.WriteString("...")
			break
		}
	}
	return text.String()
}

func getNumOfBlogs(amount int) []blog {
	count := 0
	blogList := []blog{}
	for _, v := range blogs {
		if count == amount {
			break
		}
		v.Count = count
		blogList = append(blogList, v)
		count++
	}
	return blogList
}

func getBlog(blogID int) (blog, error) {
	b := blogs[blogID]
	if b == (blog{}) {
		err := errors.New("blog id not found")
		return b, err
	}
	return b, nil
}

func readBlogDb() []blog {
	var blogList []blog
	data, _ := ioutil.ReadFile("./markdown/blog.json")
	json.Unmarshal(data, &blogList)
	return blogList
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
// 	log.Println(getBlogText(100, "./markdown/welcome.md"))
// }
