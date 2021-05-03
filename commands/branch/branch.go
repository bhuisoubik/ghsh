// Command: branch
// (c) Soubik Bhui <@soubikbhuiwk007> 2020

package branch

import (
	"context"
	"fmt"

	"github.com/google/go-github/v35/github"
	"github.com/soubikbhuiwk007/ghsh/reg"
	"github.com/soubikbhuiwk007/ghsh/vm/config"
	"golang.org/x/oauth2"
)

var token = config.AuthToken

func listBranch() []string {
	lsBranch := make([]string, 0)
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	client := github.NewClient(oauth2.NewClient(ctx, ts))
	branchList, _, err := client.Repositories.ListBranches(ctx, config.UserName, config.CurrentRepo, nil)
	if err != nil {
		fmt.Println(err)
	} else {
		for i := 0; i < len(branchList); i++ {
			lsBranch = append(lsBranch, branchList[i].GetName())
		}
	}
	return lsBranch
}

func contains(arr []string, find string) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == find {
			return true
		}
	}
	return false
}

func Branch(args []string) {
	if len(args) > 1 && config.IsInsideRepo {
		if args[1] == "change" {
			if contains(listBranch(), args[2]) {
				config.Branch = args[2]
			} else {
				fmt.Println("branch not found")
			}
		} else if args[1] == "log" {
			lsbranch := listBranch()
			for i := 0; i < len(lsbranch); i++ {
				fmt.Println(lsbranch[i])
			}
		}
	} else {
		config.PrintError("Invalid Argument || Invalid State")
	}
}

func Register() {
	reg.RegisterNewCommand("branch", Branch)
}
