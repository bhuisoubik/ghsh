// Command: rmdir
// (c) Soubik Bhui <@soubikbhuiwk007> 2020


package rmdir

import (
	"context"
	"fmt"

	"github.com/google/go-github/v35/github"
	"github.com/soubikbhuiwk007/ghsh/reg"
	"github.com/soubikbhuiwk007/ghsh/vm/config"
	"golang.org/x/oauth2"
)

var token = config.AuthToken

var filePathsDel = make([]string, 0)
var sha = make([]string, 0)

func getAllFiles(dirPath string) ([]string, []string) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	_, dirContent, _, err := client.Repositories.GetContents(ctx, config.UserName, config.CurrentRepo, dirPath, &github.RepositoryContentGetOptions{
		Ref: config.Branch,
	})
	if err != nil {
		fmt.Println(err)
	} else {
		for i := 0; i < len(dirContent); i++ {
			if *dirContent[i].Type == "file" {
				filePathsDel = append(filePathsDel, *dirContent[i].Path)
				sha = append(sha, *dirContent[i].SHA)
			} else if *dirContent[i].Type == "dir" {
				getAllFiles(*dirContent[i].Path)
			}
		}

	}

	return filePathsDel, sha
}

func Rmdir(args []string) {
	if len(args) > 1 {
		if !config.IsInsideRepo {
			ctx := context.Background()
			ts := oauth2.StaticTokenSource(
				&oauth2.Token{AccessToken: token},
			)
			client := github.NewClient(oauth2.NewClient(ctx, ts))
			_, err := client.Repositories.Delete(ctx, config.UserName, args[1])
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Repository '%v' deleted successfully", args[1])
			}
		} else {
			fmt.Println("Note: time for deleting the directory depend on the number of sub-dirs inside the folder")
			ctx := context.Background()
			ts := oauth2.StaticTokenSource(
				&oauth2.Token{AccessToken: token},
			)
			client := github.NewClient(oauth2.NewClient(ctx, ts))
			msg := "remove directory '" + args[1] + "'"
			fList, shaList := getAllFiles(config.Repo_Path + args[1])
			for i := 0; i < len(fList); i++ {
				opts := &github.RepositoryContentFileOptions{
					SHA:     &shaList[i],
					Message: &msg,
					Branch:  &config.Branch,
				}
				_, _, err := client.Repositories.DeleteFile(ctx, config.UserName, config.CurrentRepo, fList[i], opts)
				if err != nil {
					fmt.Println(err)
				}
			}
			fmt.Printf("Directory '%v' deleted successfully", args[1])
		}
	} else {
		config.PrintError("Invalid Argument")
	}
}

func Register() {
	reg.RegisterNewCommand("rmdir", Rmdir)
}
