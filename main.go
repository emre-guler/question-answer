package main

import (
	"net/http"
	"os"

	"github.com/emre-guler/question-answer/endpoints"
)

var port string = os.Getenv("PORT")

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", loginHandler)
	mux.HandleFunc("/gh-callback", ghHandler)
	http.ListenAndServe((":" + port), mux)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		endpoints.LoginPageGET(w)
	default:
		endpoints.MethodNotAllowed(w)
	}
}

func ghHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		endpoints.CallbackGET(w, r)
	default:
		endpoints.MethodNotAllowed(w)
	}
}
