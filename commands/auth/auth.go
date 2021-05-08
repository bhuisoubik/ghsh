// Command: auth
// (c) Soubik Bhui <@soubikbhuiwk007> 2021

package auth

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"syscall"

	"github.com/google/go-github/v35/github"
	"github.com/soubikbhuiwk007/ghsh/reg"
	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/oauth2"
)


func printGenTokSteps() {
	fmt.Println("Steps to Generate Personal Access Token\n\t1. Go to your Github Profile\n\t2. Go to Settings -> Developer Settings -> Personal access tokens -> Generate new token\n\t3. Copy the Token and paste it in the below prompt")
}

func setToken(token string) {
	user, _ := user.Current()
	authFname := filepath.Join(user.HomeDir, ".ghsh_auth_token")
	ioutil.WriteFile(authFname, []byte(token), 0644)
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

		setToken(token)
		fmt.Println("\nAuthentication Successful")
	}
}

func Logout() {
	user, _ := user.Current()
	authFname := filepath.Join(user.HomeDir, ".ghsh_auth_token")
	ioutil.WriteFile(authFname, []byte(" "), 0644)
	fmt.Println("Logged out successfully")
	fmt.Println("Run 'ghsh auth -login' for re-login")
}


func Auth(args []string) {
	if len(args) > 1 {
		if args[1] == "-login" {
			Login()
		} else if args[1] == "-logout" {
			Logout()
			os.Exit(0)
		} else {
			fmt.Println("Invalid Argument")
		}
	} else {
		fmt.Println("Invalid Argument")
	}
}

func Register() {
	reg.RegisterNewCommand("auth", Auth)
}