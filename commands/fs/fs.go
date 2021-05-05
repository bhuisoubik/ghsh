// Command: fs
// (c) Soubik Bhui <@soubikbhuiwk007> 2021

package fs

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/google/go-github/v35/github"
	"github.com/soubikbhuiwk007/ghsh/reg"
	"github.com/soubikbhuiwk007/ghsh/vm/config"
	"golang.org/x/oauth2"
)

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

var token = config.AuthToken

func Fs(args []string) {
	if len(args) > 2 {
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		tc := oauth2.NewClient(ctx, ts)
		client := github.NewClient(tc)

		if args[1] == "-r" && config.IsInsideRepo {
			fpath := string([]byte(config.Repo_Path)[1:]) + args[2]
			fContent, _, _, err := client.Repositories.GetContents(ctx, config.UserName, config.CurrentRepo, fpath, &github.RepositoryContentGetOptions{
				Ref: config.Branch,
			})

			if err != nil {
				fmt.Println(err)
			} else {
				content, _ := fContent.GetContent()
				fmt.Print(content)
			}
		} else if args[1] == "-w" && config.IsInsideRepo {
			sha, fpath := getSHA(args[2], config.Repo_Path) 
			msg := "Update '" + args[2] + "'"
			content, _ := ioutil.ReadFile(args[3])
			_, _, err := client.Repositories.UpdateFile(ctx, config.UserName, config.CurrentRepo, fpath, &github.RepositoryContentFileOptions{
				Message: &msg,
				Content: content,
				Branch: &config.Branch,
				SHA: &sha,
			})

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(msg)
			}
		}
	} else {
		config.PrintError("Invalid Argument")
	}
}

func Register() {
	reg.RegisterNewCommand("fs", Fs)
}