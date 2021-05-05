// Command: rls
// (c) Soubik Bhui <@soubikbhuiwk007> 2021

package rls

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

func boolEval(c []byte) bool {
	if string(c) == "Y" || string(c) == "y" {
		return true
	} 
	return false
}

func Release(args []string) {
	if len(args) > 1 {
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		tc := oauth2.NewClient(ctx, ts)
		client := github.NewClient(tc)
		rd := bufio.NewReader(os.Stdin)
		if args[1] == "new" { // rls new
			fmt.Print("Tag Name: ")
			byteTgName, _, _ := rd.ReadLine()
			tagname := string(byteTgName)
			fmt.Print("Name: ")
			byteName, _, _ := rd.ReadLine()
			name := string(byteName)
			fmt.Print("Body: ")
			body := ""
			for {
				byteNewBody, _, _ := rd.ReadLine()
				if string(byteNewBody) == "" {
					break
				} else {
					body += string(byteNewBody) + "\n"
				}
			}
			fmt.Print("Draft [Y|N]: ")
			byteDraft, _, _ := rd.ReadLine()
			draft := boolEval(byteDraft)
			fmt.Print("Pre-Release [Y|N]: ")
			bytePR, _, _ := rd.ReadLine()
			pre := boolEval(bytePR)
			rls, _, err := client.Repositories.CreateRelease(ctx, config.UserName, config.CurrentRepo, &github.RepositoryRelease{
				TagName: &tagname,
				Name: &name,
				Body: &body,
				Draft: &draft,
				Prerelease: &pre,
			})

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("New Release (tag: %v) is successfully created at %v", *rls.TagName, *rls.CreatedAt)
			}
		} else if args[1] == "latest" { // rls latest
			rls, _, err := client.Repositories.GetLatestRelease(ctx, config.UserName, config.CurrentRepo)

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("%v     -%v", *rls.TagName, *rls.Name)
			}
		} else if args[1] == "list" { // rls list
			for i := 0; i < i+1; i++ {
				rlsList, _, err := client.Repositories.ListReleases(ctx, config.UserName, config.CurrentRepo, &github.ListOptions{
					Page: i + 1,
					PerPage: 100,
				})
				if err != nil {
					fmt.Println(err)
				} else {
					if len(rlsList) > 0 {
						for li := 0; li < len(rlsList); li++ {
							fmt.Printf("\n%v -%v\n", *rlsList[li].TagName, *rlsList[li].Name)
						}
					} else {
						break
					}
				}
			}
		} else if args[1] == "get" { // rls get <tag-name>
			rls, _, err := client.Repositories.GetReleaseByTag(ctx, config.UserName, config.CurrentRepo, args[2])

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("\n%v -%v\n\n%v\n", *rls.TagName, *rls.Name, *rls.Body)
			}
		} else if args[1] == "del" { // rls del <tag-name>
			rls, _, err1 := client.Repositories.GetReleaseByTag(ctx, config.UserName, config.CurrentRepo, args[2])
			if err1 != nil {
				fmt.Println(err1)
				return
			}
			_, err2 := client.Repositories.DeleteRelease(ctx, config.UserName, config.CurrentRepo, *rls.ID)
			if err2 != nil {
				fmt.Println(err2)
			} else {
				fmt.Printf("Release (Tag: %v) is successfully deleted", args[2])
			}
		} else if args[1] == "updt" { // rls updt <flag> <tagname>
			getRls, _, err := client.Repositories.GetReleaseByTag(ctx, config.UserName, config.CurrentRepo, args[3])
			if err != nil {
				fmt.Println(err)
				return
			}
			id := *getRls.ID
			if len(args) < 4 {
				fmt.Println("Argument Not Supplied")
				return
			}
			if args[2] == "-t" { // rls updt -t <tagname>
				fmt.Print("TagName: ")
				byteTag, _, _ := rd.ReadLine()
				tagname := string(byteTag)
				_, _, err := client.Repositories.EditRelease(ctx, config.UserName, config.CurrentRepo, id, &github.RepositoryRelease{
					TagName: &tagname,
				})

				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Printf("'TagName' field is successfully updated")
				}
			} else if args[2] == "-n" { // rls updt -n <tagname>
				fmt.Print("Name: ")
				byteName, _, _ := rd.ReadLine()
				name := string(byteName)
				_, _, err := client.Repositories.EditRelease(ctx, config.UserName, config.CurrentRepo, id, &github.RepositoryRelease{
					Name: &name,
				})

				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Printf("'Name' field is successfully updated")
				}
			} else if args[2] == "-b" { // rls updt -b <tagname>
				fmt.Print("Body: ")
				body := ""
				for {
					byteNewBody, _, _ := rd.ReadLine()
					if string(byteNewBody) == "" {
						break
					} else {
						body += string(byteNewBody) + "\n"
					}
				}
				_, _, err := client.Repositories.EditRelease(ctx, config.UserName, config.CurrentRepo, id, &github.RepositoryRelease{
					Body: &body,
				})

				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Printf("'Body' field is successfully updated")
				}
			} else if args[2] == "-d" { // rls updt -d <tagname>
				fmt.Print("Draft [Y|N]: ")
				byteDraft, _, _ := rd.ReadLine()
				draft := boolEval(byteDraft)
				_, _, err := client.Repositories.EditRelease(ctx, config.UserName, config.CurrentRepo, id, &github.RepositoryRelease{
					Draft: &draft,
				})

				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Printf("'Draft' field is successfully updated")
				}
			} else if args[2] == "-p" { // rls updt -p <tagname>
				fmt.Print("Pre Release [Y|N]: ")
				bytePR, _, _ := rd.ReadLine()
				pre := boolEval(bytePR)
				_, _, err := client.Repositories.EditRelease(ctx, config.UserName, config.CurrentRepo, id, &github.RepositoryRelease{
					Prerelease: &pre,
				})

				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Printf("'Pre Release' field is successfully updated")
				}
			}
		} else if args[1] == "assets" { // rls assets <command> <argument>
			if len(args) < 3 {
				fmt.Println("Command Not Supplied")
				return
			}

			if args[2] == "list" { // rls assets list <tag-name>
				getRls, _, err := client.Repositories.GetReleaseByTag(ctx, config.UserName, config.CurrentRepo, args[3])
				if err != nil {
					fmt.Println(err)
					return
				}
				id := *getRls.ID
				for i := 0; i < i+1; i++ {
					rlsList, _, err := client.Repositories.ListReleaseAssets(ctx, config.UserName, config.CurrentRepo, id, &github.ListOptions{
						Page: i + 1,
						PerPage: 100,
					})
					if err != nil {
						fmt.Println(err)
					} else {
						if len(rlsList) > 0 {
							for li := 0; li < len(rlsList); li++ {
								fmt.Printf("\n%v - %vB\n", *rlsList[li].Name, *rlsList[li].Size)
							}
						} else {
							break
						}
					}
				}
			}
		}
	} else {
		config.PrintError("Invalid Argument")
	}
}

func Register() {
	reg.RegisterNewCommand("rls" , Release)
}