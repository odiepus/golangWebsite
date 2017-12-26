package main

import (
	"html/template"
	"log"
	"net/http"
	"fmt"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", idx)
	http.HandleFunc("/about", abut)
	http.HandleFunc("/contact", cntct)
	http.HandleFunc("/apply", aply)
	http.ListenAndServe(":8080", nil)
}

func idx(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)

	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	fmt.Println("we got here")
}
func abut(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "about.gohtml", nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	fmt.Println("we got here")
}
func cntct(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "contact.gohtml", nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	fmt.Println("we got here")
}
func aply(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "apply.gohtml", nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	fmt.Println("we got here")
}
