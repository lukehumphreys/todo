package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
)

func LoadTodos() (*Todos, error) {

	todos := &Todos{}

	// get current user
	user, err := user.Current()
	if err != nil {
		return todos, fmt.Errorf("cannot get current user: %v", err)
	}

	// set todos file path
	todos.path = fmt.Sprintf("%s/.todos", user.HomeDir)

	// check if todos file exists
	if _, err := os.Stat(todos.path); os.IsNotExist(err) {
		f, err := os.Create(todos.path)
		if err != nil {
			return todos, fmt.Errorf("cannot create todo file: %s", todos.path)
		}
		err = f.Close()
		if err != nil {
			return todos, fmt.Errorf("cannot close file: %s", todos.path)
		}
		return todos, nil
	}

	// get todos file content
	b, err := ioutil.ReadFile(todos.path)
	if err != nil {
		return todos, fmt.Errorf("cannot read todos from: %s", todos.path)
	}

	// unmarshal and return
	err = json.Unmarshal(b, todos)
	if err != nil {
		return todos, fmt.Errorf("cannot unmarshal todos from: %s", todos.path)
	}
	return todos, err
}

type Todos struct {
	path  string `json:"-"`
	Todos []Todo `json:"todos"`
}

type Todo struct {
	done    bool   `json:"-"`
	Id      int    `json:"id"`
	Message string `json:"message"`
}

func (t *Todos) List() []Todo {
	return t.Todos
}

func (t *Todos) Add(message string) {

	todo := Todo{
		Id:      len(t.Todos),
		Message: message,
	}
	t.Todos = append(t.Todos, todo)
}

func (t *Todos) Get(id int) (*Todo, error) {

	for i := range t.Todos {
		if id == t.Todos[i].Id {
			return &t.Todos[i], nil
		}
	}
	return nil, fmt.Errorf("todo %d not found", id)
}

func (t *Todos) Save(stdErr *log.Logger) {

	ts := []Todo{}
	for _, t := range t.Todos {
		if !t.done {
			t.Id = len(ts)
			ts = append(ts, t)
		}
	}
	t.Todos = ts

	b, err := json.Marshal(t)
	if err != nil {
		stdErr.Fatalf("cannot marshal todos: %v", err)
	}

	err = ioutil.WriteFile(t.path, b, 0666)
	if err != nil {
		stdErr.Fatalf("cannot write to file %s:", t.path, err)
	}
}

func (t *Todo) Done() {
	t.done = true
}
