package main

import (
	"html/template"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
)


var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Index", produtos)
	defer db.Close()
}
