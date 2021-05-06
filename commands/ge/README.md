### Command: `ge`

Edit/Update Gist

#### Usage

`ge <command> <argument>`

#### Commands

* `-d <gist-id>` to update gist description

```shell
$ ghsh: @username / > ge -d <gist-id>
```

* `-v <gist-id>` to update gist visibility (public | secret)

```shell
$ ghsh: @username / > ge -v <gist-id>
```

* `-f <gist-id>` to update gist files

```shell
$ ghsh: @username / > ge -f <gist-id>
```

***