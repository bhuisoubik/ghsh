package gc

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/google/go-github/v35/github"
	"github.com/soubikbhuiwk007/ghsh/reg"
	"github.com/soubikbhuiwk007/ghsh/vm/config"
	"golang.org/x/oauth2"
)

var token = config.AuthToken

func Gc(args []string) {
	if len(args) > 1 {
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		tc := oauth2.NewClient(ctx, ts)
		client := github.NewClient(tc)
		rd := bufio.NewReader(os.Stdin)

		if args[1] == "new" { // gc new <id>
			fmt.Print("Body: ")
			body := ""
			for {
				byteBody, _, _ := rd.ReadLine()
				if string(byteBody) == "" {
					break
				} else {
					body += string(byteBody) + "\n"
				}
			}
			cmt, _, err := client.Gists.CreateComment(ctx, args[2], &github.GistComment{
				Body: &body,
			})

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Commented Successfully created at %v", *cmt.CreatedAt)
			}
		} else if args[1] == "updt" { // gc updt <gist-id> <comment-id>
			commentID, _ := strconv.ParseInt(args[3], 10, 64) 
			fmt.Print("Body: ")
			body := ""
			for {
				byteBody, _, _ := rd.ReadLine()
				if string(byteBody) == "" {
					break
				} else {
					body += string(byteBody) + "\n"
				}
			}
			cmt, _, err := client.Gists.EditComment(ctx, args[2], commentID, &github.GistComment{
				Body: &body,
			})

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Commented Successfully updated at %v", *cmt.CreatedAt)
			}
		} else if args[1] == "del" { // gc del <gist-id> <comment-id>
			commentID, _ := strconv.ParseInt(args[3], 10, 64) 
			_, err := client.Gists.DeleteComment(ctx, args[2], commentID)

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Comment is deleted successfully")
			}
		} else if args[1] == "get" { // gc get <gist-id> <comment-id>
			commentID, _ := strconv.ParseInt(args[3], 10, 64) 
			cmt, _, err := client.Gists.GetComment(ctx, args[2], commentID)

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("\nby @%v     at: %v\n\n%v\n", *cmt.User.Login, *cmt.CreatedAt, cmt.GetBody())
			}
		} else if args[1] == "list" { // gc list <gist-id>
			for i := 0; i < i+1; i++ {
				cmnt,_, err := client.Gists.ListComments(ctx, args[2], &github.ListOptions{
					PerPage: 100,
					Page: i + 1,
				})
				if err != nil {
					fmt.Println(err)
				} else {
					if len(cmnt) > 0 {
						for li := 0; li < len(cmnt); li++ {
							fmt.Printf("\nby @%v     at: %v\n\n%v\n", *cmnt[li].User.Login, *cmnt[li].CreatedAt, cmnt[li].GetBody())	
						}
					} else {
						break
					}
				}
			}
		}

	} else {
		config.PrintError("Invalid Argument")
	}
}

func Register() {
	reg.RegisterNewCommand("gc" , Gc)
}