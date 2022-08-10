package main

import (
	"net/http"
	"os"
	"text/template"
)

var loginTemp = template.Must(template.ParseFiles("./views/login.html"))

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9091"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/login", loginHandler)
	http.ListenAndServe((":" + port), mux)
}

// TODO Login with X platform
func loginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		loginTemp.Execute(w, nil)
	case "POST":
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405: Method not allowed."))
	}
}
