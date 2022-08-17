package githubservice

import (
	"context"
	"errors"
	"log"

	"github.com/emre-guler/question-answer/models"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func GetClient(accessToken string, background context.Context) *github.Client {
	tc := oauth2.NewClient(background, oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	))
	return github.NewClient(tc)
}

func GetUserData(client *github.Client, background context.Context, accessToken string) (*models.User, error) {
	user, _, err := client.Users.Get(background, "")
	newUser := new(models.User)
	if err != nil {
		log.Println("Can't connect to Github services: ", err)
		return newUser, errors.New("can't connect to Github services")
	}
	newUser.AccessToken = accessToken
	newUser.AvatarUrl = user.GetAvatarURL()
	newUser.ProfileUrl = user.GetURL()
	newUser.FullName = user.GetName()
	newUser.Company = user.GetCompany()
	newUser.GithubId = user.GetID()
	newUser.Location = user.GetLocation()

	eturn newUser, nil
}
