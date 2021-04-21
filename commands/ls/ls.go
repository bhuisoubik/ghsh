package ls

import (
	"context"
	"fmt"

	"github.com/google/go-github/v35/github"
	"github.com/soubikbhuiwk007/ghve/reg"
	"github.com/soubikbhuiwk007/ghve/vm/config"
)

func Ls(args []string) {
	client := github.NewClient(nil)
	if args[1] == "org" {
		userid := ""
		if len(args) < 3 {
			userid = config.UserName
		} else {
			userid = args[2]
		}
		ctx := context.Background()
		orgs, _, err := client.Organizations.List(ctx, userid, nil)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("UserID: %v\n\tOrganisations:\n", userid)
			for i := 0; i < len(orgs); i++ {
				fmt.Printf("\t\t%v\n", *orgs[i].Login)
			}
		}
	} else if args[1] == "repo" {
		userid := ""
		if len(args) < 3 {
			userid = config.UserName
		} else {
			userid = args[2]
		}
		ctx := context.Background()
		opts := &github.RepositoryListOptions{Type: "public"}
		repos, _, err := client.Repositories.List(ctx, userid, opts)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("UserID: %v\n\tRepository:\n", userid)
			for i := 0; i < len(repos); i++ {
				fmt.Println("\t\t" + *repos[i].Name)
			}
		}
	} else if args[1] == "gist" {
		userid := ""
		if len(args) < 3 {
			userid = config.UserName
		} else {
			userid = args[2]
		}
		ctx := context.Background()
		gists, _, err := client.Gists.List(ctx, userid, nil)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("UserID: %v\n\tGist:\n", userid)
			for i := 0; i < len(gists); i++ {
				fmt.Printf("\t\t%v   %v  %v file(s)\n", *gists[i].Description, *gists[i].ID, len(gists[i].Files))
			}
		}
	} else {
		fmt.Println("Invalid Argument")
	}
}

func Register() {
	reg.RegisterNewCommand("ls", Ls)
}