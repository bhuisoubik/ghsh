// Command: ge
// (c) Soubik Bhui <@soubikbhuiwk007> 2021

package ge

import (
	"bufio"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/google/go-github/v35/github"
	"github.com/soubikbhuiwk007/ghsh/reg"
	"github.com/soubikbhuiwk007/ghsh/vm/config"
	"golang.org/x/oauth2"
)

var token = config.AuthToken

func Ge(args []string) {
	if len(args) > 1 {
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		tc := oauth2.NewClient(ctx, ts)
		client := github.NewClient(tc)
		rd := bufio.NewReader(os.Stdin)
		if args[1] == "-d" { // ge -d <gist-id>
			fmt.Print("Description: ")
			byteDesc, _, _ := rd.ReadLine()
			desc := string(byteDesc)
			gist, _, err := client.Gists.Edit(ctx, args[2], &github.Gist{
				Description: &desc,
			})

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Description of %v is successfully updated at %v", *gist.ID, *gist.UpdatedAt)
			}
		} else if args[1] == "-v" { // ge -v <gist-id>
			fmt.Print("Type (public | secret): ")
			byteType, _, _ := rd.ReadLine()
			var public bool = false
			if string(byteType) == "public" {
				public = true
			}
			get, _, _ := client.Gists.Get(ctx, args[2])
			gist, _, err := client.Gists.Edit(ctx, args[2], &github.Gist{
				Public: &public,
				Description: get.Description,
				Files: get.Files,
			})

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Visibility of %v is successfully updated at %v", *gist.ID, *gist.UpdatedAt)
			}
		} else if args[1] == "-f" { // ge -f <gist-id>
			gistContent := make(map[github.GistFilename]github.GistFile)
			for i := 0; i < i+1; i++ {
				fmt.Printf("Path [%v]: ", i)
				bytePath, _, _ := rd.ReadLine()
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
			gist, _, err := client.Gists.Edit(ctx, args[2], &github.Gist{
				Files: gistContent,
			})

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Files of %v is successfully updated at %v", *gist.ID, *gist.UpdatedAt)
			}
		}
	} else {
		config.PrintError("Invalid Argument")
	}
}

func Register() {
	reg.RegisterNewCommand("ge", Ge)
}