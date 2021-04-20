package pwd

import (
	"fmt"
	"github.com/soubikbhuiwk007/ghve/reg"
	"github.com/soubikbhuiwk007/ghve/vm/config"
)

func Pwd(args []string) {
	fmt.Println(config.CWD)
}

func Register() {
	reg.RegisterNewCommand("pwd", Pwd)
}