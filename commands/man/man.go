package man

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/soubikbhuiwk007/ghsh/reg"
)

func Man(args []string) {
	url := fmt.Sprintf("https://github.com/soubikbhuiwk007/ghsh/blob/master/commands/%v/README.md", args[1])
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