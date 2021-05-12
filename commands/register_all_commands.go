package commands

import (
	"github.com/soubikbhuiwk007/ghsh/commands/about"   // by: @soubikbhuiwk007 command: about
	"github.com/soubikbhuiwk007/ghsh/commands/auth"    // by: @soubikbhuiwk007 command: auth <command>
	"github.com/soubikbhuiwk007/ghsh/commands/branch"  // by: @soubikbhuiwk007 command: branch <command> <argument>
	"github.com/soubikbhuiwk007/ghsh/commands/cd"      // by: @soubikbhuiwk007 command: cd <subdirectory>
	"github.com/soubikbhuiwk007/ghsh/commands/commits" // by: @soubikbhuiwk007 command: commits <arguments>
	"github.com/soubikbhuiwk007/ghsh/commands/dir"     // by: @soubikbhuiwk007 command: dir
	"github.com/soubikbhuiwk007/ghsh/commands/exit"    // by: @soubikbhuiwk007 command: exit
	"github.com/soubikbhuiwk007/ghsh/commands/fork"    // by: @soubikbhuiwk007 command: fork <argument>
	"github.com/soubikbhuiwk007/ghsh/commands/fs"      // by: @soubikbhuiwk007 command: fs <command> <argument>
	"github.com/soubikbhuiwk007/ghsh/commands/issue"   // by: @soubikbhuiwk007 command: issue <command> <arguments>
	"github.com/soubikbhuiwk007/ghsh/commands/ls"      // by: @soubikbhuiwk007 command: ls <repo|org|gist> <userid>
	"github.com/soubikbhuiwk007/ghsh/commands/man"	   // by: @soubikbhuiwk007 command: man <command>
	"github.com/soubikbhuiwk007/ghsh/commands/mkdir"  // by: @soubikbhuiwk007 command: mkdir <> || <subdirectory>
	"github.com/soubikbhuiwk007/ghsh/commands/mkfile" // by: @soubikbhuiwk007 command: mkfile <> || <filename>
	"github.com/soubikbhuiwk007/ghsh/commands/pr"     // by: @soubikbhuiwk007 command: pr <command> <argument>
	"github.com/soubikbhuiwk007/ghsh/commands/ptoken" // by: @soubikbhuiwk007 command: get-token
	"github.com/soubikbhuiwk007/ghsh/commands/pwd"    // by: @soubikbhuiwk007 command: pwd
	"github.com/soubikbhuiwk007/ghsh/commands/readme" // by: @soubikbhuiwk007 comamnd: readme
	"github.com/soubikbhuiwk007/ghsh/commands/rls"    // by: @soubikbhuiwk007 command: rls <command> <argument>
	"github.com/soubikbhuiwk007/ghsh/commands/rmdir"  // by: @soubikbhuiwk007 command: rmdir <subdirectory|repository>
	"github.com/soubikbhuiwk007/ghsh/commands/rmfile" // by: @soubikbhuiwk007 command: rmdir <filename|gist-id>
	"github.com/soubikbhuiwk007/ghsh/commands/vsn"    // by: @soubikbhuiwk007 command: vsn
	"github.com/soubikbhuiwk007/ghsh/commands/gc" // by: @soubikbhuiwk007 command gc <command> <argument>
	"github.com/soubikbhuiwk007/ghsh/commands/ge" // by: @soubikbhuiwk007 command ge <command> <argument>
	"github.com/soubikbhuiwk007/ghsh/commands/gv" // by: @soubikbhuiwk007 command ge <argument>
)

func RegisterAll() {
	pwd.Register() // Print Working Directory
	ptoken.Register() // Print Token
	exit.Register() // Exit with Code 0
	ls.Register() // List Organisation/Repository/Gist for Any User
	cd.Register() // Change Directory
	dir.Register() // List Content inside Directory
	mkdir.Register() // Create a new Directory/Repository
	mkfile.Register() // Create a new File/Gist
	rmdir.Register() // Remove an existing Directory/Repository
	rmfile.Register() // Remove an existing File/Gist
	branch.Register() // Log/Change Branch of Repository
	fork.Register() // Fork Repository/Gist
	pr.Register() // Manage Pull Request
	vsn.Register() // Version
	auth.Register() // Authentication (login & logout)
	about.Register() // About
	commits.Register() // List Repository Commits (By Specific User)
	issue.Register() // Repo-Isuues
	rls.Register() // Manage Repo Release
	readme.Register() // View README.md of Repo
	fs.Register() // Read/Write Files
	man.Register() // Manual
	gc.Register() // Gist Comments
	ge.Register() // Gist Edit
	gv.Register() // Gist View
}
