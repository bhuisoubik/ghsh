// Command: rmfile
// (c) Soubik Bhui <@soubikbhuiwk007> 2021

package rmfile

import (
	"context"
	"fmt"

	"github.com/google/go-github/v35/github"
	"github.com/soubikbhuiwk007/ghsh/reg"
	"github.com/soubikbhuiwk007/ghsh/vm/config"
	"golang.org/x/oauth2"
)

var token = config.AuthToken

func getSHA(filename string, repoPath string) (string, string) {
	sha, fPath := "", ""
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	repoPath = string([]byte(repoPath)[1:])
	_, content, _, err := client.Repositories.GetContents(ctx, config.UserName, config.CurrentRepo, repoPath, &github.RepositoryContentGetOptions{
		Ref: config.Branch,
	})

	if err != nil {
		fmt.Println(err)
	} else {
		for i := 0; i < len(content); i++ {
			if *content[i].Name == filename {
				sha = *content[i].SHA
				fPath = *content[i].Path
				break
			}
		}
	}
	return sha, fPath
}

func Rmfile(args []string) {
	if len(args) > 1 {
		if !config.IsInsideRepo {
			ctx := context.Background()
			ts := oauth2.StaticTokenSource(
				&oauth2.Token{AccessToken: token},
			)
			tc := oauth2.NewClient(ctx, ts)
			client := github.NewClient(tc)
			_, err := client.Gists.Delete(ctx, args[1])
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Gist of ID '%v' deleted successfully", args[1])
			}
		} else {
			ctx := context.Background()
			ts := oauth2.StaticTokenSource(
				&oauth2.Token{AccessToken: token},
			)
			tc := oauth2.NewClient(ctx, ts)
			client := github.NewClient(tc)
			msg := "Delete file '" + args[1] + "'"
			sha, fp := getSHA(args[1], config.Repo_Path)
			_, _, err := client.Repositories.DeleteFile(ctx, config.UserName, config.CurrentRepo, fp, &github.RepositoryContentFileOptions{
				Message: &msg,
				SHA:     &sha,
				Branch:  &config.Branch,
			})

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("File '%v' deleted successfully", args[1])
			}
		}
	} else {
		config.PrintError("Invalid Argument")
	}
}

func Register() {
	reg.RegisterNewCommand("rmfile", Rmfile)
}
