# todo

Super simple TODO command line manager

## usage

 - `todo <TODO>`    adds <TODO>
 - `todo list`      lists TODOs
 - `todo edit <id>` edit/override TODO with specified id
 - `todo done <id>` remove TODO with specified id
 - `todo help`      prints help
 
## install

 - install [go](https://golang.org/dl/)
 - set `GOPATH`, example: `export GOPATH=$HOME/Documents/go`
 - add `$GOPATH/bin` to your path, example: `export PATH=$PATH:$GOPATH/bin`
 - clone this project to `$GOPATH/src/github.com/reisinger/todo`
 - run in that directory `go install`, now `todo` command should be available
