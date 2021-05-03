### Command: `branch`

Change/List branch of a repository

#### Usage
`branch <command> <argument>`

>  `branch` command can only be used inside a repository

* `change` command to change the current branch of a repository
```shell
$ ghsh: @username /repo/path/ (branch) > branch change <branch-name>
```
* `log` command to print all the available branches
```shell
$ ghsh: @username /repo/path/ (branch) > branch log
```
***