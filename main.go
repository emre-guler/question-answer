package main

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

var loginTemp = template.Must(template.ParseFiles("./views/login.gohtml"))

const ghReqUrl string = "https://github.com/login/oauth/authorize?scope=user:email&client_id="

var ghClientId string = os.Getenv("GITHUB_CLIENT_ID")

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "9091"
	}
	if ghClientId == "" {
		ghClientId = "ac5475d7d269e1ebaf4c"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/login", loginHandler)
	mux.HandleFunc("/gh-callback", ghHandler)
	http.ListenAndServe((":" + port), mux)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		vm := map[string]interface{}{
			"ghUrl": (ghReqUrl + ghClientId),
		}
		loginTemp.Execute(w, vm)
	case "POST":
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405: Method not allowed."))
	}
}

func ghHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("callback!")
}
