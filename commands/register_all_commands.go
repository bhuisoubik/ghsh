package commands

import (
	"github.com/soubikbhuiwk007/ghve/commands/exit" // Exit
	"github.com/soubikbhuiwk007/ghve/commands/ptoken" // PToken
	"github.com/soubikbhuiwk007/ghve/commands/pwd" // Pwd
	"github.com/soubikbhuiwk007/ghve/commands/crt" // Create
)

func RegisterAll() {
	pwd.Register() // Print Working Directory
	ptoken.Register() // Print Token
	exit.Register() // Exit with Code 0
	crt.Register() // Create New repo or gist
}