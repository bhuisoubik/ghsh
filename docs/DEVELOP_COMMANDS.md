### New Commands 

#### Requirement
* Go-sdk 1.16.x

#### Get-Source



#### Steps:

Create a folder (named as the command name) in the `/commands/` directory.

Inside the folder create at least two files

* README (that contains the doc of the command)
* A Go source file (it should be named as the `command-name`)

##### The Go Source File

```go
package commandName

import (
    "github.com/soubikbhuiwk007/ghsh/reg"
)

// the `args` parameter gives the arguments passed in the shell 
func CommandNameFunc(args []string) {
    ....
}

func Register() {
    reg.RegisterNewCommand("commandName", CommandNameFunc)
}
```

##### Edit the `commands/register_all_commands.go`
```go
package commands

import (
    ...
    "github.com/soubikbhuiwk007/ghgs/commands/commandName" // by: @author command: commandName          
)

func RegisterAll() {
    ...
    commandName.Register() // A small description
}
```

:zap: And Done !!! :zap:

Also try to make a [pull request](../CONTRIBUTING.md)

***