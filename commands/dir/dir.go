package dir

import (
	"context"
	"fmt"

	"github.com/google/go-github/v35/github"
	"github.com/soubikbhuiwk007/ghve/reg"
	"github.com/soubikbhuiwk007/ghve/vm/config"
	"golang.org/x/oauth2"
)

var token = config.AuthToken

func Dir(args []string) {
	if config.IsInsideRepo {
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		client := github.NewClient(oauth2.NewClient(ctx, ts))
		_, content, _, err := client.Repositories.GetContents(ctx, config.UserName, config.CurrentRepo, config.Repo_Path, &github.RepositoryContentGetOptions{
			Ref: config.Branch,
		})
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print("Type\tName\n\n")
		for i := 0; i < len(content); i++ {
			fmt.Printf("%v\t%v\n", *content[i].Type, *content[i].Name)
		}
	} else {
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		client := github.NewClient(oauth2.NewClient(ctx, ts))
		repoList, _, _ := client.Repositories.List(ctx, "", nil)
		gistList, _, _ := client.Gists.List(ctx, "", nil)

		fmt.Print("Type\tName/ID\n\n")
		for i := 0; i < len(repoList); i++ {
			fmt.Printf("repo\t%v\n", *repoList[i].Name)
		}

		for i := 0; i < len(gistList); i++ {
			fmt.Printf("gist\t%v\t\t%v\n", *gistList[i].ID, *gistList[i].Description)
		}
	}
}

func Register() {
	reg.RegisterNewCommand("dir", Dir)
}