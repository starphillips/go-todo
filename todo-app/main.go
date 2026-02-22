package main

import "time"

type Todo struct {
	ID          int
	Title       string
	Completed   bool
	CompletedAt time.Time
	CreatedAt   time.Time
}

type Todos []Todo

func (todos *Todos) add(title string) {
	todo := Todo{
		ID:          len(*todos) + 1,
		Title:       title,
		Completed:   false,
		CompletedAt: time.Time{},
		CreatedAt:   time.Now(),
	}

	*todos = append(*todos, todo)
}

// delete
func (todos *Todos) delete(id int) {

	for i := 0; i < len(*todos); i++ {
		// if id == Todo.ID { }
		// that is the struct var

	}
	// check IDs
	// if ID matches the one being requested to be deleted on the command line, then delete it
	// if not check the next ID
}

// complete
func (todos *Todos) complete(id int) {
	// check ids
	// if ID matches the one being requested to be completed on the command line, then delete it
	// if not check the next ID
}

// display all tasks

// subcommands to call these funcs

// https://www.janekbieser.dev/posts/cli-app-with-subcommands-in-go/
