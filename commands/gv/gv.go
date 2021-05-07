package gv

import (
	"context"
	"fmt"

	"github.com/google/go-github/v35/github"
	"github.com/soubikbhuiwk007/ghsh/reg"
	"github.com/soubikbhuiwk007/ghsh/vm/config"
	"golang.org/x/oauth2"
)

var token = config.AuthToken

func Gv(args []string) {
	if len(args) > 1 {
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		tc := oauth2.NewClient(ctx, ts)
		client := github.NewClient(tc)

		gist, _, err := client.Gists.Get(ctx, args[1])

		if err != nil {
			fmt.Println(err)
		} else {
			for f, c := range gist.Files {
				fmt.Printf("\nFileName: %v\n\n", f)
				fmt.Println(c.GetContent())
			}
		}
	} else {
		config.PrintError("Invalid Argument")
	}
}

func Register() {
	reg.RegisterNewCommand("gv", Gv)
}