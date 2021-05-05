// Command: readme
// (c) Soubik Bhui <@soubikbhuiwk007> 2021

package readme

import (
	"context"
	"fmt"

	"github.com/google/go-github/v35/github"
	"github.com/soubikbhuiwk007/ghsh/reg"
	"github.com/soubikbhuiwk007/ghsh/vm/config"
	"golang.org/x/oauth2"
)

var token = config.AuthToken

func Readme(args []string) {
	if config.IsInsideRepo {
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		tc := oauth2.NewClient(ctx, ts)
		client := github.NewClient(tc)

		repoContent, _, err := client.Repositories.GetReadme(ctx, config.UserName, config.CurrentRepo, &github.RepositoryContentGetOptions{
			Ref: config.Branch,
		})

		if err != nil {
			fmt.Println(err)
		} else {
			content, _ := repoContent.GetContent()
			fmt.Print(content)
		}
	} else {
		config.PrintError("Invalid State")
	}
}

func Register() {
	reg.RegisterNewCommand("readme", Readme)
}