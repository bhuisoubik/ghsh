package mkfile

import (
	"bufio"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/google/go-github/v35/github"
	"github.com/soubikbhuiwk007/ghve/reg"
	"github.com/soubikbhuiwk007/ghve/vm/config"
	"golang.org/x/oauth2"
)

var token = config.AuthToken

func gistConfig() *github.Gist {
	rd := bufio.NewReader(os.Stdin)
	fmt.Print("Description: ")
	byteDesc, _, _ := rd.ReadLine()
	desc := string(byteDesc)
	fmt.Print("Type (public | secret): ")
	byteType, _, _ := rd.ReadLine()
	var public bool = false
	if string(byteType) == "public" {
		public = true
	}
	gistContent := make(map[github.GistFilename]github.GistFile)
	for i := 0; i < i + 1; i++ {
		fmt.Printf("Path [%v]: ", i)
		bytePath,_,_ := rd.ReadLine()
		if string(bytePath) == "" {
			break
		} else {
			fileByteContent, err := ioutil.ReadFile(string(bytePath))
			if err != nil {
				fmt.Println(err)
			} else {				
				fileContent := string(fileByteContent)
				filename := github.GistFilename(filepath.Base(string(bytePath)))
				file := github.GistFile{
					Content: &fileContent,
				}
				gistContent[filename] = file
			}
		}
	}
	return &github.Gist{
		Description: &desc,
		Public: &public,
		Files: gistContent,
	}
}

func Mkfile(args []string) {
	if !config.IsInsideRepo {
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		client := github.NewClient(oauth2.NewClient(ctx, ts))
		gist,_, err := client.Gists.Create(ctx, gistConfig())
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Gist Successfully Created (ID:", *gist.ID, ")")
		}
	} else {
		if len(args) > 1 {
			ctx := context.Background()
			ts := oauth2.StaticTokenSource(
				&oauth2.Token{AccessToken: token},
			)
			client := github.NewClient(oauth2.NewClient(ctx, ts))
			filePath := string([]byte(config.Repo_Path)[1:]) + args[1]
			commitMsg := "create file " + args[1]
			opts := github.RepositoryContentFileOptions{
				Message: &commitMsg,
				Content: []byte(" "),
			}
			_, _, err := client.Repositories.CreateFile(ctx, config.UserName, config.CurrentRepo, filePath, &opts)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("File '%v' is created successfully\n", args[1])
			}
		} else {
			fmt.Println("Invalid Argument")
		}
	}
}

func Register() {
	reg.RegisterNewCommand("mkfile", Mkfile)
}