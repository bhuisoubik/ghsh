package cd

import (
	"fmt"

	"github.com/soubikbhuiwk007/ghve/reg"
	"github.com/soubikbhuiwk007/ghve/vm/config"
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
	c_path = string([]rune(c_path)[:len([]rune(c_path))-1])
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

func changeDirInsideRepo() {
	
}

func Cd(args []string) {
	if args[1] != ".." {
		if contains(getRepo(), args[1]) && !config.IsInsideRepo {
			config.CWD += args[1] + "/"
			config.IsInsideRepo = true
			config.CurrentRepo = args[1]
		} else if config.IsInsideRepo {
			changeDirInsideRepo()
		} else {
			fmt.Printf("'%v' no such directory\n", args[1])
		}
	} else {
		if config.CWD == "/" {
			config.IsInsideRepo = false
			config.CurrentRepo = ""
		} else {
			config.CWD = movPathBack(config.CWD)
			config.IsInsideRepo =  true
		}
	}
}

func Register() {
	reg.RegisterNewCommand("cd",Cd)
}