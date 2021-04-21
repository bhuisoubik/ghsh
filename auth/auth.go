package auth

import (
	"context"
	"fmt"
	"io/ioutil"
	"syscall"

	"github.com/google/go-github/v35/github"
	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/oauth2"
)

func printGenTokSteps() {
	fmt.Println("Steps to Generate Personal Access Token\n\t1. Go to your Github Profile\n\t2. Go to Settings -> Developer Settings -> Personal access tokens -> Generate new token\n\t3. Copy the Token and paste it in the below prompt")
}

func Login() {
	printGenTokSteps()
	fmt.Print("GitHub Token: ")
	byteToken, _ := terminal.ReadPassword(int(syscall.Stdin))

	token := string(byteToken)
	if token == "" {
		fmt.Println("Prompt cannot be empty")
	} else {
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		tc := oauth2.NewClient(ctx, ts)
	
		client := github.NewClient(tc)
	
		_, _, err := client.Users.Get(ctx, "")
		if err != nil {
			fmt.Printf("\nerror: %v\n", err)
			fmt.Println("Authentication Failed")
			return
		}
		
		ioutil.WriteFile("auth/.gh_access_token", []byte(token), 0)
		fmt.Println("\nAuthentication Successful")
	}
}

func Logout() {
	ioutil.WriteFile("auth/.gh_access_token", []byte(""), 0)
	fmt.Println("Logged out successfully")
	fmt.Println("Run 'ghve auth -login' for re-login")
}