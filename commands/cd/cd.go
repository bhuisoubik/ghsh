package cd

import (
	"context"
	"strings"
	"fmt"

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
var repo_path string = "/"

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
	_, dirContent, _, err  := client.Repositories.GetContents(ctx, config.UserName, getRepoNameFromPath(config.CWD), path, nil)
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
		} else if config.IsInsideRepo {
			if contains(getDirContent(repo_path), args[1]) {
				config.CWD += args[1] + "/"
				repo_path += args[1] + "/"
			} else {
				fmt.Printf("'%v' no such directory\n", args[1])
			}
		} else {
			fmt.Printf("'%v' no such directory\n", args[1])
		}
	} else if config.IsInsideRepo && args[1] == ".." {
		config.CWD = movPathBack(config.CWD)
		repo_path = movPathBack(repo_path)
		if config.CWD == "/" {
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