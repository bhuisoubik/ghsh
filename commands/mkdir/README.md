### Command: `mkdir`

Use the `mkdir` command to create a Directory/Repository

#### Usage

`mkdir <argument>`

* If the command is used outside a repository then it will prompt to create a new repository. No argument needs to be supplied.
```shell
$ ghsh: @username / > mkdir
```
```
Repository Name:
Description: 
Type (private | public):
HomePage URL:
```

* If the command is used inside a repository then it will create a new directory. The name of directory needs to be supplied as an argument.
```shell
$ ghsh: @username /repo/path/ (branch) > mkdir <dir-name>
```
***