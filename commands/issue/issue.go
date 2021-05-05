// Command: fork
// (c) Soubik Bhui <@soubikbhuiwk007> 2021

package issue

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/google/go-github/v35/github"
	"github.com/soubikbhuiwk007/ghsh/reg"
	"github.com/soubikbhuiwk007/ghsh/vm/config"
	"golang.org/x/oauth2"
)

var token = config.AuthToken

func Issue(args []string) {
	if len(args) > 1 {
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		tc := oauth2.NewClient(ctx, ts)
		client := github.NewClient(tc)
		issueInt,_ := strconv.Atoi(args[len(args)-1])
		if args[1] == "close" {  // issue close <issue-number>
			stt := "closed"
			issue, _, err := client.Issues.Edit(ctx, config.UserName, config.CurrentRepo, issueInt, &github.IssueRequest{
				State: &stt,
			})

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Issue (Id: %v | Number: %v) closed successfully at %v", *issue.ID, *issue.Number, issue.ClosedAt)
			}
		} else if args[1] == "open" { // issue open <issue-number>
			stt := "open"
			issue, _, err := client.Issues.Edit(ctx, config.UserName, config.CurrentRepo, issueInt, &github.IssueRequest{
				State: &stt,
			})

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Issue (Id: %v | Number: %v) re-opened successfully", *issue.ID, *issue.Number)
			}
		} else if args[1] == "get" { // issue get <issue-number>
			issue, _, err := client.Issues.Get(ctx, config.UserName, config.CurrentRepo, issueInt)

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("\n#%v    %v\n\n%v\n\nCreated At: %v\n", *issue.Number, *issue.Title, *issue.Body, *issue.CreatedAt)
			}
		} else if args[1] == "state" { // issue state <issue-number>
			issue, _, err := client.Issues.Get(ctx, config.UserName, config.CurrentRepo, issueInt)

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("State (Issue #%v): %v", *issue.Number, *issue.State)
			}
		} else if args[1] == "list" {  // issue list <state(open|closed|all)>
			for i := 0; i < i+1; i++ {
				issueList, _, err := client.Issues.ListByRepo(ctx, config.UserName, config.CurrentRepo, &github.IssueListByRepoOptions{
					State: args[2],
					ListOptions: github.ListOptions{
						PerPage: 100,
						Page: i + 1,
					},
				})
				if err != nil {
					fmt.Println(err)
				} else {
					if len(issueList) > 0 {
						for li := 0; li < len(issueList); li++ {
							fmt.Printf("\n#%v    %v\n\n%v\n\nCreated At: %v\n", *issueList[li].Number, *issueList[li].Title, *issueList[li].Body, *issueList[li].CreatedAt)
						}
					} else {
						break
					}
				}
			}
		} else if args[1] == "updt" {
			rd := bufio.NewReader(os.Stdin)
			if args[2] == "-t" { // issue updt -t <issue-number>
				fmt.Print("Title: ")
				byteTitle, _, _ := rd.ReadLine()
				title := string(byteTitle)
				issue, _, err := client.Issues.Edit(ctx, config.UserName, config.CurrentRepo, issueInt, &github.IssueRequest{
					Title: &title,
				})

				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Printf("'Title' field is updated at %v", *issue.UpdatedAt)
				}
			} else if args[2] == "-b" { // issue updt -b <issue-number>
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
				issue, _, err := client.Issues.Edit(ctx, config.UserName, config.CurrentRepo, issueInt, &github.IssueRequest{
					Body: &body,
				})

				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Printf("'Body' field is updated at %v", *issue.UpdatedAt)
				}
			} else if args[2] == "-l" { // issue updt -l <issue-number>
				fmt.Print("Label (seperate by ','): ")
				byteLabels, _, _ := rd.ReadLine()
				labels := strings.Split(string(byteLabels), ",")

				issue, _, err := client.Issues.Edit(ctx, config.UserName, config.CurrentRepo, issueInt, &github.IssueRequest{
					Labels: &labels,
				})

				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Printf("'Label' field is updated at %v", *issue.UpdatedAt)
				}
			} else if args[2] == "-a" { // issue updt -a <issue-number>
				fmt.Print("Assignees (seperate by ','): ")
				byteAssign, _, _ := rd.ReadLine()
				assignees := strings.Split(string(byteAssign), ",")

				issue, _, err := client.Issues.Edit(ctx, config.UserName, config.CurrentRepo, issueInt, &github.IssueRequest{
					Assignees: &assignees,
				})

				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Printf("'Assignees' field is updated at %v", *issue.UpdatedAt)
				}
			}
		} else if args[1] == "new" { // issue new
			rd := bufio.NewReader(os.Stdin)
			fmt.Print("Title: ")
			byteTitle, _, _ := rd.ReadLine()
			title := string(byteTitle)
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
			fmt.Print("Label (seperate by ','): ")
			byteLabels, _, _ := rd.ReadLine()
			labels := strings.Split(string(byteLabels), ",")
			fmt.Print("Assignees (seperate by ','): ")
			byteAssign, _, _ := rd.ReadLine()
			assignees := strings.Split(string(byteAssign), ",")

			issue, _, err := client.Issues.Create(ctx, config.UserName, config.CurrentRepo, &github.IssueRequest{
				Title: &title,
				Body: &body,
				Labels: &labels,
				Assignees: &assignees,
			})

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("\nNew is issue (Number: %v | Id: %v) is created at %v\n", *issue.Number, *issue.ID, *issue.CreatedAt)
			}
		} else if args[1] == "lock" { // issue lock <issue-number>
			rd := bufio.NewReader(os.Stdin)
			fmt.Print("Lock Reason [off-topic, too heated, resolved, spam] : ")
			byteLR, _, _ := rd.ReadLine()
			lr := string(byteLR)
			_, err := client.Issues.Lock(ctx, config.UserName, config.CurrentRepo, issueInt, &github.LockIssueOptions{
				LockReason: lr,
			})

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Issue #%v is locked", issueInt)
			}
		} else if args[1] == "comment" { 
			if args[2] == "del" { // issue comment del <comment-id>
				commentID, _ := strconv.ParseInt(args[3], 10, 64)
				_, err := client.Issues.DeleteComment(ctx, config.UserName, config.CurrentRepo,commentID)

				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("Comment is successfully deleted")
				}
			} else if args[2] == "get" { // issue comment get <comment-id>
				commentID, _ := strconv.ParseInt(args[3], 10, 64)
				comment, _, err := client.Issues.GetComment(ctx, config.UserName, config.CurrentRepo, commentID)

				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Printf("\n@%v        -%v\n\n%v\n", *comment.User.Login, *comment.CreatedAt, *comment.Body)
				}
			} else if args[2] == "new" { // issue comment new <issue-number>
				rd := bufio.NewReader(os.Stdin)
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
				issueCmnt, _, err := client.Issues.CreateComment(ctx, config.UserName, config.CurrentRepo, issueInt, &github.IssueComment{
					Body: &body,
				})

				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Printf("Comment (ID: %v) is created at %v by @%v", *issueCmnt.ID, *issueCmnt.CreatedAt, *issueCmnt.User.Login)
				}
			} else if args[2] == "list" { // issue comment list <issue-number>
				for i := 0; i < i+1; i++ {
					issCmtList, _, err := client.Issues.ListComments(ctx, config.UserName, config.CurrentRepo, issueInt, &github.IssueListCommentsOptions{
						ListOptions: github.ListOptions{
							PerPage: 100,
							Page: i + 1,
						},
					})
					if err != nil {
						fmt.Println(err)
					} else {
						if len(issCmtList) > 0 {
							for li := 0; li < len(issCmtList); li++ {
								fmt.Printf("\n@%v        -%v\n\n%v\n", *issCmtList[li].User.Login, *issCmtList[li].CreatedAt, *issCmtList[li].Body)
							}
						} else {
							break
						}
					}
				}
			} else if args[2] == "updt" { // issue comment updt <comment-id>
				rd := bufio.NewReader(os.Stdin)
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
				commentID, _ := strconv.ParseInt(args[3], 10, 64)
				issueCmnt, _, err := client.Issues.EditComment(ctx, config.UserName, config.CurrentRepo, commentID, &github.IssueComment{
					Body: &body,
				})

				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Printf("Comment (ID: %v) is updated at %v by @%v", *issueCmnt.ID, *issueCmnt.CreatedAt, *issueCmnt.User.Login)
				}
			}
		}
	} else {
		config.PrintError("Invalid Argument")
	}
}

func Register() {
	reg.RegisterNewCommand("issue", Issue)
}