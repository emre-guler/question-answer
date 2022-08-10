package main

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

var loginTemp = template.Must(template.ParseFiles("./views/login.gohtml"))

var port string = os.Getenv("PORT")
var ghReqUrl string = os.Getenv("GITHUB_AUTH_REQUEST_URL")

type LoginVM struct {
	ghReqUrl string
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", loginHandler)
	mux.HandleFunc("/gh-callback", ghHandler)
	http.ListenAndServe((":" + port), mux)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		var loginVM = LoginVM{ghReqUrl}
		loginTemp.Execute(w, loginVM)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405: Method not allowed."))
	}
}

func ghHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("callback!")
	switch r.Method {
	case "GET":
		// TODO callback
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405: Method not allowed."))
	}
}
