// Command: pr
// (c) Soubik Bhui <@soubikbhuiwk007> 2020

package pr

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

func getPR(base string) *github.NewPullRequest {
	fmt.Println("Note: Leave the title prompt if you want add Issue Number and vice-versa")
	var title, head, body string
	rd := bufio.NewReader(os.Stdin)
	fmt.Print("Title: ")
	t, _, _ := rd.ReadLine()
	title = string(t)
	fmt.Print("Issue Number: ")
	i, _, _ := rd.ReadLine()
	issueNum, _ := strconv.Atoi(string(i))
	fmt.Println("Body: ")
	body = ""
	for {
		b, _, _ := rd.ReadLine()
		if string(b) == "" {
			break
		} else {
			body += string(b) + "\n"
		}
	}
	head = config.UserName + ":" + config.Branch
	if title == "" {
		return &github.NewPullRequest{
			Issue: &issueNum,
			Head:  &head,
			Body:  &body,
			Base:  &base,
		}
	}
	return &github.NewPullRequest{
		Title: &title,
		Head:  &head,
		Body:  &body,
		Base:  &base,
	}
}

func PullRequest(args []string) {
	if len(args) > 1 {
		prInt, _ := strconv.Atoi(args[len(args)-1])
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		tc := oauth2.NewClient(ctx, ts)
		client := github.NewClient(tc)
		if args[1] == "new" && config.IsInsideRepo { // New Pull Request
			pr, _, err := client.PullRequests.Create(ctx, config.UserName, config.CurrentRepo, getPR(args[2]))
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("New Pull Request is created! (Number: %v | Id: %v)", *pr.Number, *pr.ID)
			}
		} else if args[1] == "list" && config.IsInsideRepo { // List Pull Request
			if len(args) > 2 {
				for i := 0; i < i+1; i++ {
					prList, _, err := client.PullRequests.List(ctx, config.UserName, config.CurrentRepo, &github.PullRequestListOptions{
						State: args[2],
						ListOptions: github.ListOptions{
							PerPage: 100,
							Page:    i + 1,
						},
					})
					if err != nil {
						fmt.Println(err)
					} else {
						if len(prList) > 0 {
							for li := 0; li < len(prList); li++ {
								fmt.Printf("\n#%v   %v\n\n\tBase.Ref: %v\n\tHead.Ref: %v\n\nCreated At: %v\n", *prList[li].Number, *prList[li].Title, *prList[li].Base.Ref, *prList[li].Head.Ref, *prList[i].CreatedAt)
							}
						} else {
							break
						}
					}
				}
			}
		} else if args[1] == "close" && config.IsInsideRepo { // Close A Pull Request
			prInt, _ := strconv.Atoi(args[2])
			stt := "closed"
			pr, _, err := client.PullRequests.Edit(ctx, config.UserName, config.CurrentRepo, prInt, &github.PullRequest{
				State: &stt,
			})
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("%v is closed successfully at %v\n", *pr.Number, pr.GetClosedAt())
			}
		} else if args[1] == "open" && config.IsInsideRepo { // Open a Pull Request
			prInt, _ := strconv.Atoi(args[2])
			stt := "open"
			pr, _, err := client.PullRequests.Edit(ctx, config.UserName, config.CurrentRepo, prInt, &github.PullRequest{
				State: &stt,
			})
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("%v is re-opened successfully at %v\n", *pr.Number, *pr.UpdatedAt)
			}
		} else if args[1] == "get" && config.IsInsideRepo { // Get a specific Pull Request
			if len(args) > 2 {
				prInt, _ := strconv.Atoi(args[2])
				pr, _, err := client.PullRequests.Get(ctx, config.UserName, config.CurrentRepo, prInt)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Printf("\n#%v   %v\n\n\tBase.Ref: %v\n\tHead.Ref: %v\n\nBody: %v\n\nCreated At: %v\n", *pr.Number, *pr.Title, *pr.Base.Ref, *pr.Head.Ref, *pr.Body, *pr.CreatedAt)
				}
			}
		} else if args[1] == "state" && config.IsInsideRepo { // State of a Pull Request ('open' or 'close')
			pr, _, err := client.PullRequests.Get(ctx, config.UserName, config.CurrentRepo, prInt)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("PR-Number: %v\nState: %v\n", *pr.Number, *pr.State)
			}
		} else if args[1] == "merge" && config.IsInsideRepo { // Merge a Pull Request
			rd := bufio.NewReader(os.Stdin)
			fmt.Print("Commit Message: ")
			byteCommitMsg, _, _ := rd.ReadLine()
			prMergeRes, _, err := client.PullRequests.Merge(ctx, config.UserName, config.CurrentRepo, prInt, string(byteCommitMsg), nil)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Pull-Request (Number: %v) is successfully merged\nSHA: %v\n", prInt, prMergeRes.GetSHA())
			}
		} else if args[1] == "updt" && config.IsInsideRepo { // Update a Pull Request
			rd := bufio.NewReader(os.Stdin)
			if args[2] == "-t" { // Update Title
				fmt.Print("New Title: ")
				byteNewTitle, _, _ := rd.ReadLine()
				title := string(byteNewTitle)
				prEditted,_,err := client.PullRequests.Edit(ctx, config.UserName, config.CurrentRepo, prInt, &github.PullRequest{
					Title: &title,
				})

				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Printf("'Title' field is successfully edited at %v\n", *prEditted.UpdatedAt)
				}
			} else if args[2] == "-b" { // Update Body
				rd := bufio.NewReader(os.Stdin)
				fmt.Print("New Body: ")
				body := ""
				for {
					byteNewBody, _, _ := rd.ReadLine()
					if string(byteNewBody) == "" {
						break
					} else {
						body += string(byteNewBody) + "\n"
					}
				}
				prEdit, _, err := client.PullRequests.Edit(ctx, config.UserName, config.CurrentRepo, prInt, &github.PullRequest{
					Body: &body,
				})

				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Printf("'Body' field is successfully edited at %v\n", *prEdit.UpdatedAt)
				}
			} else if args[2] == "-m" { // Update MaintainerCanModify
				rd := bufio.NewReader(os.Stdin)
				fmt.Print("Maintainer Can Modify ([Y]es | [N]o): ")
				byteMcm, _, _ := rd.ReadLine()		
				mcm := false
				if string(byteMcm) == "Y" || string(byteMcm) == "y"	{
					mcm = true
				}
				prEdit, _, err := client.PullRequests.Edit(ctx, config.UserName, config.CurrentRepo, prInt, &github.PullRequest{
					MaintainerCanModify: &mcm,
				})

				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Printf("'MaintainerCanModify' field is successfully edited at %v\n", *prEdit.UpdatedAt)
				}
			} else if args[2] == "-br" { // Update Base.Ref
				rd := bufio.NewReader(os.Stdin)
				fmt.Print("Base Reference: ")
				byteBr, _, _ := rd.ReadLine()
				br := string(byteBr)

				prEdit, _, err := client.PullRequests.Edit(ctx, config.UserName, config.CurrentRepo, prInt, &github.PullRequest{
					Base: &github.PullRequestBranch{
						Ref: &br,
					},
				})

				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Printf("'Base.ref' field is successfully edited at %v\n", *prEdit.UpdatedAt)
				}
			}
		} else if args[1] == "diff" && config.IsInsideRepo { // Print the Differences
			pr, _, err := client.PullRequests.Get(ctx, config.UserName, config.CurrentRepo, prInt)

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Commits: %v\nAdditions: %v\nDeletions: %v\nChanged Files: %v", *pr.Commits, *pr.Additions, *pr.Deletions, *pr.ChangedFiles)
			}
		} else if args[1] == "commits" && config.IsInsideRepo{ // Print all the Commits
			for i := 0; i < i+1; i++ {
				prListCm, _, err := client.PullRequests.ListCommits(ctx, config.UserName, config.CurrentRepo, prInt ,&github.ListOptions{
					PerPage: 100,
					Page:    i + 1,
				})

				if err != nil {
					fmt.Println(err)
				} else {
					if len(prListCm) > 0 {
						for li := 0; li < len(prListCm); li++ {
							fmt.Printf("\nBy:%v     CommitMsg:%v\n", *prListCm[li].Author.Login, *prListCm[li].Commit.Message)
						}
					} else {
						break
					}
				}
			}
		} else { // Not a Valid Command
			config.PrintError("Invalid Argument")
		}
	} else { // No Argument Passed
		config.PrintError("Invalid Argument")
	}
}

func Register() {
	reg.RegisterNewCommand("pr", PullRequest)
}
