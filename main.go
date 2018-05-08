package main

import(
	"html/template"
    "net/http"
    "log"
)

func handler(w http.ResponseWriter, r *http.Request){
	tmpl, _ := template.ParseFiles("./html/home.html", "./html/default.html")
	tmpl.ExecuteTemplate(w, "layout", nil)
}


func main(){
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

    http.HandleFunc("/", handler)
    log.Println("Listening...")
    http.ListenAndServe(":8080", nil)
}
