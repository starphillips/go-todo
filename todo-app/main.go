package main

import (
	"fmt"
	"strconv"
	"time"
)

// the type for one task (blueprint)
type Todo struct {
	ID          int
	Title       string
	Completed   bool
	CompletedAt time.Time
	CreatedAt   time.Time
}

// the type for a slice of tasks
type Todos []Todo

// td - the variable holding the list - refers to the actual instance of your todo list - the actual list you're modifying
func (td *Todos) add(title string) {
	item := Todo{ //  the variable holding one new task
		ID:          len(*td) + 1,
		Title:       title,
		Completed:   false,
		CompletedAt: time.Time{},
		CreatedAt:   time.Now(),
	}

	*td = append(*td, item)
}

// delete
func (td *Todos) delete(id int) {
	for i := 0; i < len(*td); {
		if item.ID == id {

		}

	}
	// check IDs
	// if ID matches the one being requested to be deleted on the command line, then delete it
	// if not check the next ID
}

// complete
func (td *Todos) complete(id int) {
	// check ids
	// if ID matches the one being requested to be completed on the command line, then delete it
	// if not check the next ID
}

// func display all tasks

// subcommands to call these funcs

var td Todos

type command struct {
	Name string
	Help string
	Run  func(args []string) error
}

var commands = []command{
	{Name: "add", Help: "add an item", Run: addCmd(&td)},
	{Name: "delete", Help: "delete an item", Run: deleteCmd(&td)},
	{Name: "done", Help: "mark an item as done", Run: completeCmd(&td)},
}

func addCmd(td *Todos) func(args []string) error {
	return func(args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("missing title")
		}
		title := args[0]
		td.add(title)
		return nil
	}
	// parse args to get the title of the task to be added
	// call the add func with the title
}

func deleteCmd(td *Todos) func(args []string) error {
	return func(args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("missing id")
		}
		id := args[0]
		integer, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println("failed to convert string to integer", err)
		}
		td.delete(integer)
		return nil
	}
	// parse args to get the ID of the task to be deleted
	// call the delete func with the ID

}

func completeCmd(td *Todos) func(args []string) error {
	return func(args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("missing id")
		}
		done := args[0]
		integer, err := strconv.Atoi(done)
		if err != nil {
			fmt.Println("failted to convert string to integer", err)
		}
		td.complete(integer)
		return nil

	}
	// parse args to get the ID of the task to be completed
	// call the complete func with the ID
}
