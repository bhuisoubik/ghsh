package config

import (
	"context"
	"fmt"
	"io/ioutil"
	"os/user"
	"path/filepath"

	"github.com/google/go-github/v35/github"
	"golang.org/x/oauth2"
)

func getuserid() *string {
	token := getAuthToken()
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	
	client := github.NewClient(oauth2.NewClient(ctx, ts))
	user, _, err := client.Users.Get(ctx, "")
	var un string
	if err != nil {
		un = ""
	} else {
		un = *user.Login
	}
	return &un
}

func getAuthToken() string {
	user, _ := user.Current()
	authFname := filepath.Join(user.HomeDir, ".ghsh_auth_token")
	byteToken, _ := ioutil.ReadFile(authFname)
	if len(byteToken) < 1 {
		return ""
	}
	return string(byteToken)
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

func PrintError(msg string) {
	fmt.Println("Error: "+ msg)
}

var (
	CWD string = "/"
	UserName string = *getuserid()
	IsInsideRepo bool = false
	CurrentRepo string = ""
	AuthToken string = getAuthToken()
	Repo_Path string = "/"
	Branch string = ""
)