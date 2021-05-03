### Command: `dir`

It is used to print the contents of a directory

#### Usage
* If used outside a repository
```shell
$ ghsh: @username / > dir
```
```
Type    Name/ID
repo    <repo-name>
gist    <gist-id>       <gist-description>
```
* If used inside a repository
```shell
$ ghsh: @username /repo/path/ (branch-name) > dir
```
```
Type    Name
file    <file-name>
dir     <dir-name>
```
***