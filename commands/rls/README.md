### Command: `rls`

Manage Release of current Repository

#### Usage
`rls <command> <argument>`

#### Commands

* `new` to create a new release

```shell
$ ghsh: @username repo/path/ (branch) > rls new
```

* `latest` to get latest release

```shell
$ ghsh: @username repo/path/ (branch) > rls latest
```

* `list` to list all the release

```shell
$ ghsh: @username repo/path/ (branch) > rls list
```

* `get <tag-name>` to get a specific release.

```shell
$ ghsh: @username repo/path/ (branch) > rls get <tag-name>
```

* `del` to delete a repository

```shell
$ ghsh: @username repo/path/ (branch) > rls del <tag-name>
```

* `updt <command> <tag-name>` to update a release

    * `-t <tag-name>` to update tagname
    * `-n <tag-name>` to update name
    * `-b <tag-name>` to update body
    * `-d <tag-name>` to change draft value (yes | no)
    * `-p <tag-name>` to change pre-release value (yes | no)

```shell
$ ghsh: @username repo/path/ (branch) > rls updt <command> <tag-name>
```

* `assets list <tag-name>` to list assets for a release

```shell
$ ghsh: @username repo/path/ (branch) > rls assets list <tag-name>
```

***