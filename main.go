package main

import(
	"html/template"
    "net/http"
    "log"
)

type blog struct{
	Title string
}

func handler(w http.ResponseWriter, r *http.Request){
	tmpl, _ := template.ParseFiles("./html/layout.html", "./html/home.html")
	tmpl.ExecuteTemplate(w, "layout", nil)
}

func handlerBlog(w http.ResponseWriter, r *http.Request){
	b := blog{}
	b.Title = r.URL.Path[len("/blog/"):]
	log.Printf("title: %s", b.Title)
	tmpl, _ := template.ParseFiles("./html/layout.html", "./html/blog.html")
	tmpl.ExecuteTemplate(w, "layout", b)
}

func main(){
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/blog/", handlerBlog)
    http.HandleFunc("/", handler)
    log.Println("Listening...")
    http.ListenAndServe(":8080", nil)
}
