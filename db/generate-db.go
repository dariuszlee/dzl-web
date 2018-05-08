package main

import (
    // "bufio"
	"encoding/json"
    "io/ioutil"
	"fmt"
	"strings"
)

type Blog struct {
	Name string
	Desc string
	ImgPath string
}

func main(){
	f, _ := ioutil.ReadFile("./full.txt")

	pos := 0
	posNext := 0
	for i:= 0; i <= 10; i++ {
		var ImgBuilder strings.Builder
		fmt.Fprintf(&ImgBuilder, "./img%d", i)

		posNext += 10
		var NameBuilder strings.Builder
		fmt.Fprintf(&NameBuilder, "log%d - ", i)
		NameBuilder.Write(f[pos:posNext])
		pos += 10

		posNext += 10
		var DescBuilder strings.Builder
		DescBuilder.Write(f[pos:posNext])
		posNext += 10

		blog := Blog{NameBuilder.String(), DescBuilder.String(), ImgBuilder.String()}
		j, _ := json.Marshal(blog)
		fmt.Print(string(j))
	}
}
