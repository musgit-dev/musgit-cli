# Musgit CLI - Git-like management of your music practices in your terminal

## Overview

Musgit is a tool for managing your music practices.

Musgit provides:

* Adding various music pieces
* Practicing and evaluating these pieces
* Working with lessons


## Installation

`musgit-cli` can be installed:

```go
go install github.com/musgit-dev/musgit-cli@latest
```


## Usage

* Start new lesson

```
musgit-cli start lesson
```

* Add new piece

```
musgit-cli pieces add --name "Piece Name" --composer "Composer" --complexity 1
```

* Practice  piece #1

```
musgit-cli pieces practice 1
```

## Configuration

All your practices are stored in a Sqlite database (located in `$HOME/.musgit/musgit.db`)

If you want to store results somewhere else use one of the following:

* Adjust `db-uri` parameter in `$HOME/.musgit/musgit.yaml` file.
* Pass DB file path to `--db-uri` CLI flag.
* Add `MUSGIT_DB_URI` ENV variable.


## License

Musgit is released under the Apache 2.0 license. See [LICENSE.txt](LICENSE.txt).

