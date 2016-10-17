# todo

Super simple TODO command line manager

## usage

 - `todo <TODO>`         adds <TODO>
 - `todo list`           lists TODOs
 - `todo edit <id>`      edit/override TODO
 - `todo swap <id> <id>` swaps specified TODOs
 - `todo done <id>`      remove TODO from the list
 - `todo pop`            removes first TODO (same as 'todo done 0')
 - `todo help`           prints help
 
## install

 - install [go](https://golang.org/dl/)
 - set `GOPATH`, example: `export GOPATH=$HOME/Documents/go`
 - add `$GOPATH/bin` to your path, example: `export PATH=$PATH:$GOPATH/bin`
 - clone this project to `$GOPATH/src/github.com/reisinger/todo`
 - run in that directory `go install`, now `todo` command should be available
