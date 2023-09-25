package main

import (
	"html/template"
	"net/http"
)

func greet(w http.ResponseWriter, r *http.Request) {
  // fmt.Fprintf(w, "Hello World ! dingooooooooooooo %s", time.Now())

  tmpl := template.Must(template.ParseFiles("test.html"))
  tmpl.Execute(w, nil)
  
}

func test(w http.ResponseWriter, r *http.Request) {
  
}

func new(w http.ResponseWriter, r *http.Request) {
  tmpl := template.Must(template.ParseFiles("new.html"))
  tmpl.Execute(w, nil)
  
}

func old(w http.ResponseWriter, r *http.Request) {
  tmpl := template.Must(template.ParseFiles("old.html"))
  tmpl.Execute(w, nil)
  
}

func main() {
  http.HandleFunc("/", greet)
  http.HandleFunc("/fade_out_demo", test )
  http.HandleFunc("/new", new )
  http.HandleFunc("/old", old )
  http.ListenAndServe(":8080", nil)
}
