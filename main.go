package main

import (
	"html/template"
	"log"
	"net/http"
)

type blog struct {
	Title       string
	Description string
}

type project struct {
	Title       string
	Description string
}

type homepage struct {
	Projects []project
	Blogs    []blog
}

func handlerHome(w http.ResponseWriter, r *http.Request) {
	blogs := getBlogs()
	projs := getProjects()
	homeData := homepage{projs, blogs}
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
