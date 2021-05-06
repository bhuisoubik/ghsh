### Command: `gc`

Manage Gist Comment

#### Usage

`gc <command> <argument>...`

#### Commands

* `new <gist-id>` to create a new gist comment

```shell
$ ghsh: @username / > gc new <gist-id>
```

* `updt <gist-id> <comment-id>` to update a gist comment

```shell
$ ghsh: @username / > gc updt <gist-id> <comment-id>
```

* `del <gist-id> <comment-id>` to delete a gist comment

```shell
$ ghsh: @username / > gc del <gist-id> <comment-id>
```

* `get <gist-id> <comment-id>` to get a specific gist comment

```shell
$ ghsh: @username / > gc get <gist-id> <comment-id>
```

* `list <gist-id>` to list comments for a gist

```shell
$ ghsh: @username / > gc list <gist-id>
```

***