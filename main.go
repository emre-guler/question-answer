package main

import (
	"net/http"
	"os"
	"text/template"
)

var loginTemp = template.Must(template.ParseFiles("./view/login.html"))

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9091"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/login", loginHandler)
	http.ListenAndServe((":" + port), mux)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
	case "POST":
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
