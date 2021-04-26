package cd

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/go-github/v35/github"
	"github.com/soubikbhuiwk007/ghve/reg"
	"github.com/soubikbhuiwk007/ghve/vm/config"
	"golang.org/x/oauth2"
)

func contains(arr []string, find string) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == find {
			return true
		}
	}
	return false
}

var token = config.AuthToken

func getAllRepo() []string {
	repos := make([]string, 0)
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

func movPathBack(c_path string) string {
	if c_path != "/"{
		c_path = string([]rune(c_path)[:len([]rune(c_path))-1])
	}
	for i := len(c_path) - 1; i >= 0; i-- {
		if c_path[i] != '/' {
			ru := []rune(c_path)
			c_path = string(ru[:len(ru)-1])
		} else {
			break
		}
	}

	return c_path
}

func getRepoNameFromPath(path string) string {
	return strings.Split(path, "/")[1]
}

func getDirContent(path string) []string{
	dirs := make([]string, 0)
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	_, dirContent, _, err  := client.Repositories.GetContents(ctx, config.UserName, getRepoNameFromPath(config.CWD), path, &github.RepositoryContentGetOptions{
		Ref: config.Branch,
	})
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(dirContent); i++ {
		if *dirContent[i].Type == "dir" {
			dirs = append(dirs, *dirContent[i].Name)
		}
	}
	return dirs
}

func Cd(args []string) {
	if args[1] != ".." {
		if contains(getAllRepo(), args[1]) && !config.IsInsideRepo {
			config.CWD += args[1] + "/"
			config.IsInsideRepo = true
			config.CurrentRepo = args[1]
			config.Branch = config.GetFirstBranch(args[1])
		} else if config.IsInsideRepo {
			if contains(getDirContent(config.Repo_Path), args[1]) {
				config.CWD += args[1] + "/"
				config.Repo_Path += args[1] + "/"
			} else {
				fmt.Printf("'%v' no such directory\n", args[1])
			}
		} else {
			fmt.Printf("'%v' no such directory\n", args[1])
		}
	} else if config.IsInsideRepo && args[1] == ".." {
		config.CWD = movPathBack(config.CWD)
		config.Repo_Path = movPathBack(config.Repo_Path)
		if config.CWD == "/" {
			config.Branch = ""
			config.IsInsideRepo = false
			config.CurrentRepo = ""
		} else {
			config.IsInsideRepo =  true
		}
	}
}

func Register() {
	reg.RegisterNewCommand("cd",Cd)
}