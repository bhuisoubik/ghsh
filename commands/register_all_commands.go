package commands

import (
	"github.com/soubikbhuiwk007/ghve/commands/exit" // Exit
	"github.com/soubikbhuiwk007/ghve/commands/ptoken" // PToken
	"github.com/soubikbhuiwk007/ghve/commands/pwd" // Pwd
	"github.com/soubikbhuiwk007/ghve/commands/ls" // List
)

func RegisterAll() {
	pwd.Register() // Print Working Directory
	ptoken.Register() // Print Token
	exit.Register() // Exit with Code 0
	ls.Register() // List Organistion/Repository/Gist for Any User
}