package main

import (
	"html/template"
	"net/http"
	"strings"
	"fmt"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8888", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ip := r.Header.Get("X-Real-IP")
	//ip := r.RemoteAddr
	t, _ := template.ParseFiles("view.tmpl")
	fmt.Println(ip)
	t.Execute(w, ip)
}
