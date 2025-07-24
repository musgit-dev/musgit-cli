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

## License

Musgit is released under the Apache 2.0 license. See [LICENSE.txt](LICENSE.txt).

