package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/emre-guler/question-answer/models"
	githubservice "github.com/emre-guler/question-answer/service"
	"github.com/gin-gonic/gin"
)

var githubClientSecret string = os.Getenv("GITHUB_CLIENT_SECRET")
var githubRequestUrl string = os.Getenv("GITHUB_AUTH_REQUEST_URL")
var githubClientId string = os.Getenv("GITHUB_CLIENT_ID")
var background = context.Background()

const githubAccesTokenUrl string = "https://github.com/login/oauth/access_token"

func LoginGetHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var loginVm = models.LoginPageViewModel{GithubRequestUrl: githubRequestUrl + githubClientId}
		ctx.HTML(http.StatusOK, "login.gohtml", gin.H{
			"GithubRequestUrl": loginVm,
		})
	}
}

func CallbackGetHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		errParam, err := ctx.Params.Get("error")
		if errParam != "" {
			log.Println("Github callback error: ", err)
			ctx.Redirect(http.StatusMovedPermanently, "/login")
			return
		}

		codeParam, _ := ctx.Params.Get("code")
		reqValues := url.Values{
			"client_id":     {githubClientId},
			"client_secret": {githubClientSecret},
			"code":          {codeParam},
			"accept":        {"json"}}
		req, _ := http.NewRequest("POST", githubAccesTokenUrl, strings.NewReader(reqValues.Encode()))
		req.Header.Set(
			"Accept", "application/json")

		res, resErr := http.DefaultClient.Do(req)
		if resErr != nil {
			log.Println("Github token request error: ", resErr)
			ctx.Redirect(http.StatusMovedPermanently, "/login")
			return
		}

		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			log.Println("Github token request error: ", res.StatusCode)
			ctx.Redirect(http.StatusMovedPermanently, "/login")
		}

		var access models.Access

		if err := json.NewDecoder(res.Body).Decode(&access); err != nil {
			log.Println("Json decode error: ", err)
			ctx.Redirect(http.StatusMovedPermanently, "/login")
			return
		}

		if access.Scope != "read:user" {
			log.Println("Wrong authority, scope error: ", access.Scope)
			ctx.Redirect(http.StatusMovedPermanently, "/login")
			return
		}

		var client = githubservice.GetClient(access.AccessToken, background)
		userData, dataErr := githubservice.GetUserData(client, background, access.AccessToken)
		if dataErr != nil {
			// Loglama zaten yukarıdaki servisin içerisinde yapılıyor.
			ctx.Redirect(http.StatusMovedPermanently, "/login")
			return
		}

		fmt.Println(userData)
	}
}
