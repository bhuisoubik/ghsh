### Command: `rmdir`

Delete a directory/repository

#### Usage
`rmdir <argument>`

* If the command is used outside a repository then it will delete an repository. The `repository-name` must be supplied as an argument
```shell
$ ghsh: @username / > rmdir <repo-name>
```
```
Repository '<repo-name>' deleted successfully
```

* If the command is used inside a repository then it will delete
the specified folder. The `dir-name` should be supplied as an argument
```shell
$ ghsh: @username /repo/path/ (branch) > rmdir <dir-name>
```
```
Note: time for deleting the directory depend on the number of sub-dirs inside the folder
Directory '<dir-name>' deleted successfully
```

***
