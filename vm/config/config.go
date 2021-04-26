package config

import (
	"context"
	"io/ioutil"

	"github.com/google/go-github/v35/github"
	"golang.org/x/oauth2"
)

func getuserid() **string {
	byteToken, _ := ioutil.ReadFile("auth/.gh_access_token")
	token := string(byteToken)
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	
	client := github.NewClient(oauth2.NewClient(ctx, ts))
	user, _, _ := client.Users.Get(ctx, "")
	return &user.Login
}

func getAuthToken() string {
	byteToken, _ := ioutil.ReadFile("auth/.gh_access_token")
	token := string(byteToken)
	return token
}

func GetFirstBranch(repo string) string {
	br := ""
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: AuthToken},
	)
	client := github.NewClient(oauth2.NewClient(ctx, ts))
	repository,_, _ := client.Repositories.Get(ctx, UserName, repo)
	br = *repository.DefaultBranch
	return br
}

var (
	CWD string = "/"
	UserName string = **getuserid()
	IsInsideRepo bool = false
	CurrentRepo string
	AuthToken string = getAuthToken()
	Repo_Path string = "/"
	Branch string = ""
)