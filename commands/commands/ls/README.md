### Command: `ls`

Use the `ls` command to list **public** `repository`/`organisation`/`gist` for an `user`

#### Usage
`ls <command> <argument>`

* `repo <user-id>` command to list all the **public** repositories for an user
```shell
$ ghsh: @username / > ls repo <user-id>
```
```
UserID: <user-id>
        Repository:
                repo_name..
```
* `gist <user-id>` command to list all the **public** gists for an user
```shell
$ ghsh: @username / > ls gist <user-id>
```
```
UserID: <user-id>
        Gist:
                <gist-id>       <file-count>
```
* `org <user-id>` command to list the **public** organisations for an user
```shell
$ ghsh: @username / > ls org <user-id>
```
```
UserID: <user-id>
        Organisations:
                org_name..
```

***