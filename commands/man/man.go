package man

import (
	"os/exec"
	"runtime"

	"github.com/soubikbhuiwk007/ghsh/reg"
)

func Man(args []string) {
	url := "https://github.com/soubikbhuiwk007/ghsh/blob/master/docs/COMMANDS.md#" + args[1]
	switch runtime.GOOS {
	case "windows" :
		exec.Command("cmd", "/c", "start", url).Start()
	case "darwin":
		exec.Command("open", url).Start()
	default:
		exec.Command("xdg-open", url).Start()
	}
}

func Register() {
	reg.RegisterNewCommand("man", Man)
}