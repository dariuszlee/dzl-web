package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type project struct {
	Title       string
	Description string
	Id          int
}

type blogPage struct {
	Blogs []blog
}

type homepage struct {
	ActiveProject project
	Projects      []project
	ActiveBlog    blog
	Blogs         []blog
}

func handlerHome(w http.ResponseWriter, r *http.Request) {
	slider := 3
	blogs := getNumOfBlogs(slider)
	projs := getProjects(slider)
	homeData := homepage{projs[0], projs[1:], blogs[0], blogs[1:]}
	tmpl, _ := template.ParseFiles("./html/layout.html", "./html/home.html")
	tmpl.ExecuteTemplate(w, "layout", homeData)
}

func handlerBlog(w http.ResponseWriter, r *http.Request) {
	blogIdStr := r.URL.Path[len("/blog/"):]
	blogId, _ := strconv.Atoi(blogIdStr)
	blogData, _ := getBlog(blogId)
	blog := formatBlog(blogData.Path)
	tmpl, _ := template.ParseFiles("./html/layout.html", "./html/blog.html")
	tmpl.Parse(blog)
	tmpl.ExecuteTemplate(w, "layout", blogData)
}

func handlerAbout(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("./html/layout.html", "./html/about.html")
	tmpl.ExecuteTemplate(w, "layout", nil)
}

func handlerBlogs(w http.ResponseWriter, r *http.Request) {
	pageSize := 10
	var blogs blogPage
	blogs.Blogs = getNumOfBlogs(pageSize)
	log.Println(blogs.Blogs)
	tmpl, _ := template.ParseFiles("./html/layout.html", "./html/blogs.html")
	tmpl.ExecuteTemplate(w, "layout", blogs)
}

func handlerProjects(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("./html/layout.html", "./html/projects.html")
	tmpl.ExecuteTemplate(w, "layout", nil)
}

func main() {
	// Initialize
	getBlogs()

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/blog/", handlerBlog)
	http.HandleFunc("/about/", handlerAbout)
	http.HandleFunc("/blogs/", handlerBlogs)
	http.HandleFunc("/projects/", handlerProjects)
	http.HandleFunc("/", handlerHome)
	log.Println("Listening...")

	var isDebug = flag.Bool("d", true, "Set debug mode")
	flag.Parse()

	if *isDebug {
		log.Panic(http.ListenAndServe(":8081", nil))
	} else {
		log.Panic(http.ListenAndServe(":8080", nil))
	}
}
