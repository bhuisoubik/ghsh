package commits

import (
	"context"
	"fmt"

	"github.com/google/go-github/v35/github"
	"github.com/soubikbhuiwk007/ghsh/reg"
	"github.com/soubikbhuiwk007/ghsh/vm/config"
	"golang.org/x/oauth2"
)

var token = config.AuthToken

func Commits(args []string) {
	if len(args) > 1 {
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		tc := oauth2.NewClient(ctx, ts)
		client := github.NewClient(tc)
		author := ""
		if len(args) < 3 {
			author = config.UserName
		} else {
			author = args[2]
		}
		for i := 0; i < i+1; i++ {
			cmtList, _, err := client.Repositories.ListCommits(ctx, config.UserName, config.CurrentRepo, &github.CommitsListOptions{
				Author: author,
				ListOptions: github.ListOptions{
					PerPage: 100,
					Page:    i + 1,
				},
			})
			if err != nil {
				fmt.Println(err)
			} else {
				if len(cmtList) > 0 {
					for li := 0; li < len(cmtList); li++ {
						fmt.Printf("\nBy:%v     CommitMsg:%v\n", *cmtList[li].Author.Login, *cmtList[li].Commit.Message)
					}
				} else {
					break
				}
			}
		}
	} else {
		config.PrintError("Invalid Argument")
	}
}

func Register() {
	reg.RegisterNewCommand("commits", Commits)
}
