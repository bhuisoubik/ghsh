// Command: mkdir
// (c) Soubik Bhui <@soubikbhuiwk007> 2020

package mkdir

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v35/github"
	"github.com/soubikbhuiwk007/ghsh/reg"
	"github.com/soubikbhuiwk007/ghsh/vm/config"
	"golang.org/x/oauth2"
)

var token = config.AuthToken

func repo() *github.Repository {
	rd := bufio.NewReader(os.Stdin)
	fmt.Print("Repository Name: ")
	name, _, _ := rd.ReadLine()
	repoName := string(name)
	fmt.Print("Description: ")
	byteDesc, _, _ := rd.ReadLine()
	desc := string(byteDesc)
	fmt.Print("Type (private | public): ")
	vis, _, _ := rd.ReadLine()
	mode := string(vis)
	var private bool = true
	if mode == "public" {
		private = false
	}
	fmt.Print("HomePage URL: ")
	homeUrl, _ := rd.ReadString('\n')
	repo := &github.Repository{
		Name:        &repoName,
		Description: &desc,
		Private:     &private,
		Homepage:    &homeUrl,
	}
	return repo
}

func Mkdir(args []string) {
	if !config.IsInsideRepo {
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		client := github.NewClient(oauth2.NewClient(ctx, ts))
		newRepo, _, err := client.Repositories.Create(ctx, "", repo())
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("New Repository is Created: " + *newRepo.FullName)
			fmt.Println("Visit https://www.github.com/" + config.UserName + "/" + *newRepo.Name)
		}
	} else {
		if len(args) > 1 {
			ctx := context.Background()
			ts := oauth2.StaticTokenSource(
				&oauth2.Token{AccessToken: token},
			)
			client := github.NewClient(oauth2.NewClient(ctx, ts))
			folderPath := string([]byte(config.Repo_Path)[1:]) + args[1] + "/.gitkeep"
			msg := "create directory " + args[1]
			opts := github.RepositoryContentFileOptions{
				Message: &msg,
				Content: []byte(" "),
				Branch:  &config.Branch,
			}
			_, _, err := client.Repositories.CreateFile(ctx, config.UserName, config.CurrentRepo, folderPath, &opts)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Directory '%v' is created successfully\n", args[1])
			}
		} else {
			config.PrintError("Invalid Argument")
		}
	}
}

func Register() {
	reg.RegisterNewCommand("mkdir", Mkdir)
}
