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

var (
	CWD string = "/"
	UserName string = **getuserid()
	IsInsideRepo bool = false
	CurrentRepo string
	AuthToken string = getAuthToken()
	Repo_Path string = "/"
)