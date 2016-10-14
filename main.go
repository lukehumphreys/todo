package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var options = make(map[string]action)
var stdErr *log.Logger
var stdOut *log.Logger

type action func(*Todos, []string) ([]string, error)

func init() {

	options["list"] = list
	options["edit"] = edit
	options["done"] = done
	options["help"] = help

	stdErr = log.New(os.Stderr, "", 0)
	stdOut = log.New(os.Stdout, "", 0)
}

func main() {

	args := os.Args[1:]
	todos, err := LoadTodos()
	if err != nil {
		stdErr.Fatal(err)
	}
	defer todos.Save(stdErr)

	// no args passed
	if len(args) == 0 {
		out, err := help(todos, args)
		print(out, err)
		return
	}

	// check available option
	if f, ok := options[args[0]]; ok {
		out, err := f(todos, args[1:])
		print(out, err)
		return
	}

	// no args: add message to TODOs
	print(add(todos, args))
}

func print(lines []string, err error) {

	if err != nil {
		msg := Red(fmt.Sprintf("error: %v\n", err))
		stdErr.Fatal(msg)
	}

	for _, line := range lines {
		stdOut.Println(line)
	}
}

func add(todos *Todos, args []string) ([]string, error) {

	t := strings.Join(args, " ")
	todos.Add(t)
	return []string{"ok"}, nil
}

func list(todos *Todos, args []string) ([]string, error) {

	lines := []string{}
	for _, todo := range todos.List() {
		id := Blue(fmt.Sprintf("[%d] ", todo.Id))
		msg := Green(todo.Message)
		lines = append(lines, id + msg)
	}
	return lines, nil
}

func edit(todos *Todos, args []string) ([]string, error) {

	id, err := getId(args)
	if err != nil {
		return []string{}, err
	}

	todo, err := todos.Get(id)
	if err != nil {
		return []string{}, err
	}

	todo.Message = strings.Join(args[1:], " ")
	return []string{"ok"}, nil
}

func done(todos *Todos, args []string) ([]string, error) {

	id, err := getId(args)
	if err != nil {
		return []string{}, err
	}

	todo, err := todos.Get(id)
	if err != nil {
		return []string{}, err
	}

	todo.Done()
	return []string{"ok"}, err
}

func help(*Todos, []string) ([]string, error) {

	lines := []string{
		Green(">> super simple TODO list manager <<"),
		"",
		Blue("  todo <TODO>    ") + "adds <TODO>",
		Blue("  todo list      ") + "lists TODOs",
		Blue("  todo edit <id> ") + "edit/override TODO with specified id",
		Blue("  todo done <id> ") + "remove TODO with specified id",
		Blue("  todo help      ") + "prints this message",
		"",
	}
	return lines, nil
}

func getId(args []string) (int, error) {

	if len(args) == 0 {
		return -1, fmt.Errorf("command needs todo index number as argument")
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return -1, fmt.Errorf("invalid todo index %s", args[0])
	}
	return id, nil
}
