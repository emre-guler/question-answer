package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/emre-guler/question-answer/globals"
	"github.com/emre-guler/question-answer/models"
	"github.com/emre-guler/question-answer/service/dbservice"
	"github.com/emre-guler/question-answer/service/githubservice"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var githubClientSecret string = os.Getenv("GITHUB_CLIENT_SECRET")
var githubRequestUrl string = os.Getenv("GITHUB_AUTH_REQUEST_URL")
var githubClientId string = os.Getenv("GITHUB_CLIENT_ID")
var background = context.Background()

const githubAccesTokenUrl string = "https://github.com/login/oauth/access_token"

func LoginGetHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		isLoggedIn := IsLoggedIn(ctx)
		if isLoggedIn {
			ctx.Redirect(http.StatusMovedPermanently, "/app")
		}
		var githubRequestUrl string = githubRequestUrl + githubClientId
		ctx.HTML(http.StatusOK, "login.gohtml", gin.H{
			"GithubRequestUrl": githubRequestUrl,
		})
	}
}

func CallbackGetHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		isLoggedIn := IsLoggedIn(ctx)
		if isLoggedIn {
			ctx.Redirect(http.StatusMovedPermanently, "/app")
		}
		errParam := ctx.Query("error")
		if errParam != "" {
			log.Println("Github callback error.")
			ctx.Redirect(http.StatusMovedPermanently, "/login")
			return
		}

		codeParam := ctx.Query("code")
		if codeParam == "" {
			log.Println("Github callback error: ")
			ctx.Redirect(http.StatusMovedPermanently, "/login")
			return
		}

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
			return
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
		isUserExist := dbservice.IsUserExist(userData.GithubId)
		if !isUserExist {
			// Bu user daha önce kayıt olmamış ise
			result := dbservice.InsertUser(userData)
			if !result {
				// Loglama zaten yukarıdaki servisin içerisinde yapılıyor.
				ctx.Redirect(http.StatusMovedPermanently, "/login")
				return
			}
		} else {
			updateResult := dbservice.UpdateUserAccessToken(userData)
			if !updateResult {
				// Loglama zaten yukarıdaki servisin içerisinde yapılıyor.
				ctx.Redirect(http.StatusMovedPermanently, "/login")
				return
			}
		}

		// TODO Login işlemleri...
		session := sessions.Default(ctx)
		session.Set(globals.Userkey, userData.GithubId)
		session.Save()

		ctx.Redirect(http.StatusMovedPermanently, "/app")
	}
}

func IsLoggedIn(ctx *gin.Context) bool {
	session := sessions.Default(ctx)
	user := session.Get(globals.Userkey)
	return user != nil
}
