package fork

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/go-github/v35/github"
	"github.com/soubikbhuiwk007/ghve/reg"
	"github.com/soubikbhuiwk007/ghve/vm/config"
	"golang.org/x/oauth2"
)

var token = config.AuthToken

func Fork(args []string) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	client := github.NewClient(oauth2.NewClient(ctx, ts))
	if len(args) > 1 {
		if args[1] == "repo" {
			repo, _, err := client.Repositories.CreateFork(ctx, strings.Split(args[2],"/")[0], strings.Split(args[2],"/")[1], nil)
			if err != nil {
				fmt.Println(err)
				fmt.Printf("Check out https://www.github.com/%v\n", repo.FullName)
			}
		} else if args[1] == "gist" {
			gistFork, _, err := client.Gists.Fork(ctx, args[2])
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Gist Forked Successfully: %v", *gistFork.HTMLURL)
			}
		}
	} else {
		fmt.Println("Invalid Argument")
	}
}

func Register() {
	reg.RegisterNewCommand("fork", Fork)
}