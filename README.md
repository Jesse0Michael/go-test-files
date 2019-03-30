## Go-Test-Files

Go test files is a repository for media files I use to test that I build a CLI around. Get a url to any media file available in this repo, quickly, by using keywords to filter that type of file you need. The url will be copied to your clipboard for convenience.

## Installation

```
go get -u github.com/jesse0michael/go-test-files/gofiles
```

## Usage

usage: `gofiles [keywords...]`

```
gofiles audio 6channel
```

This command will filter down to a file that can be identified as being audio with 6 channels.

When multiple files match a filter, a random file will be chosen.

If no arguments are provide any file will be selected at random, unless a `GOFILES_DEFAULT` enviornment variable is set.
