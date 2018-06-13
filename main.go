package main

import (
	"html/template"
	"log"
	"net/http"
)

type blog struct {
	Title       string
	Description string
	Id          int
}

type project struct {
	Title       string
	Description string
	Id          int
}

type homepage struct {
	ActiveProject project
	Projects      []project
	ActiveBlog    blog
	Blogs         []blog
}

func handlerHome(w http.ResponseWriter, r *http.Request) {
	slider := 5
	blogs := getBlogs(slider)
	projs := getProjects(slider)
	homeData := homepage{projs[0], projs[1:], blogs[0], blogs[1:]}
	tmpl, _ := template.ParseFiles("./html/layout.html", "./html/home.html")
	tmpl.ExecuteTemplate(w, "layout", homeData)
}

func handlerBlog(w http.ResponseWriter, r *http.Request) {
	b := blog{}
	b.Title = r.URL.Path[len("/blog/"):]
	blog := formatBlog("./markdown/welcome.md")
	tmpl, _ := template.ParseFiles("./html/layout.html", "./html/blog.html")
	tmpl.Parse(blog)
	tmpl.ExecuteTemplate(w, "layout", b)
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/blog/", handlerBlog)
	http.HandleFunc("/", handlerHome)
	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}
