package endpoints

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var loginTemp = template.Must(template.ParseFiles("./views/login.gohtml"))

var port string = os.Getenv("PORT")
var githubRequestUrl string = os.Getenv("GITHUB_AUTH_REQUEST_URL")
var githubClientId string = os.Getenv("GITHUB_CLIENT_ID")
var githubClientSecret string = os.Getenv("GITHUB_CLIENT_SECRET")
var background = context.Background()

type LoginVM struct {
	GithubRequestUrl string
}

type Access struct {
	AccessToken string `json:"access_token"`
	Scope       string
}

func MethodNotAllowed(w http.ResponseWriter) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("405: Method not allowed."))
}

func LoginPageGET(w http.ResponseWriter) {
	var loginVM = LoginVM{githubRequestUrl + githubClientId}
	loginTemp.Execute(w, loginVM)
}

func RedirectToLogin(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, ("http://localhost:" + port + "/login"), http.StatusSeeOther)
}

func CallbackGET(w http.ResponseWriter, r *http.Request) {
	cbErr := r.URL.Query().Get("error")
	if cbErr != "" {
		RedirectToLogin(w, r)
		return
	}
	cbCode := r.URL.Query().Get("code")
	reqValues := url.Values{"client_id": {githubClientId}, "client_secret": {githubClientSecret}, "code": {cbCode}, "accept": {"json"}}
	req, _ := http.NewRequest("POST", "https://github.com/login/oauth/access_token", strings.NewReader(reqValues.Encode()))
	req.Header.Set(
		"Accept", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Print(err)
		RedirectToLogin(w, r)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Println("Access token failed: ", res.StatusCode)
		RedirectToLogin(w, r)
		return
	}
	var access Access

	if err := json.NewDecoder(res.Body).Decode(&access); err != nil {
		log.Println("JSON err: ", err)
		RedirectToLogin(w, r)
		return
	}

	if access.Scope != "read:user" {
		log.Println("Wrong token scope: ", access.Scope)
		RedirectToLogin(w, r)
		return
	}
	client := getGithubClient(access.AccessToken)
	user, _, err := client.Users.Get(background, "")
	if err != nil {
		fmt.Println("Can't connect to GitHub services.")
		RedirectToLogin(w, r)
		return
	}
	// TODO register user or login.
	fmt.Println(user)
}

func getGithubClient(accessToken string) *github.Client {
	tc := oauth2.NewClient(background, oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	))
	return github.NewClient(tc)
}
