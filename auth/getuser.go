package auth

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v35/github"
	"golang.org/x/oauth2"
)

func GetUser() *github.User {
	var us *github.User;
	token := os.Getenv("GHVE_USER_AUTH_TOKEN")
	if token == "" {
		fmt.Println("GHVE has not been authorised")
	} else {
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		tc := oauth2.NewClient(ctx, ts)
	
		client := github.NewClient(tc)
	
		user, _, err := client.Users.Get(ctx, "")
		if err != nil {
			fmt.Printf("\nerror: %v\n", err)
			fmt.Println("Invalid Token")
			return nil
		} else {
			us = user
		}		
	}
	return us
}