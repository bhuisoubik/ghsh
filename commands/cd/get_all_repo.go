package cd

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/google/go-github/v35/github"
	"golang.org/x/oauth2"
)

func getRepo() []string {
	repos := make([]string, 0)
	btyeToken, _ := ioutil.ReadFile("auth/.gh_access_token")
	token := string(btyeToken)
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	repoList, _, err := client.Repositories.List(ctx, "", nil)

	if err != nil {
		fmt.Println(err)
	} else {
		for i := 0; i < len(repoList); i++ {
			repos = append(repos, *repoList[i].Name)
		}
	}
	return repos
}