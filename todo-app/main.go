package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/rodaine/table"
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
	for i := 0; i < len(*td); i++ {
		if (*td)[i].ID == id {
			newList := append((*td)[:i], (*td)[i+1:]...) // Creates a new slcie without i (the deleted task)
			*td = newList
			fmt.Printf("Item deleted\n")
			return
		}
	}
	fmt.Printf("There is no matching ID to delete\n")
	// check IDs
	// if ID matches the one being requested to be deleted on the command line, then delete it
	// if not check the next ID
}

// complete
func (td *Todos) complete(id int) {
	for i := 0; i < len(*td); i++ {
		if (*td)[i].ID == id {
			(*td)[i].Completed = true
			(*td)[i].CompletedAt = time.Now()
			fmt.Printf("Item marked as complete\n")
			return
		}
	}
	fmt.Printf("There is no matching ID to complete\n")
	// check ids
	// if ID matches the one being requested to be completed on the command line, then cahnge its complete to completed
	// if not check the next ID
}

// func display all tasks
func (td *Todos) showList() {
	tbl := table.New("ID", "Title", "Completed?", "Created At", "Completed At")

	for _, item := range *td {
		tbl.AddRow(
			item.ID,
			item.Title,
			item.Completed,
			item.CreatedAt.Format(time.RFC822),
			item.CompletedAt.Format(time.RFC822),
		)

	}
	tbl.Print()
}

// subcommands to call these funcs

var td Todos // we create the real task list here

type command struct {
	Name string
	Help string
	Run  func(args []string) error // Function takes in one param and returns an error e.g. Run will be the function we create, that takes in 
}

var commands = []command{ // we create a list of commands here - this is made from a slice of the struct
	{Name: "add", Help: "add an item", Run: addCmd(&td)}, // ensure each follow the structs format
	{Name: "delete", Help: "delete an item", Run: deleteCmd(&td)},
	{Name: "done", Help: "mark an item as done", Run: completeCmd(&td)},
	{Name: "list", Help: "shows all items on your todo list", Run: listCmd(&td)},
}

func addCmd(td *Todos) func(args []string) error {
	return func(args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("missing title\n")
		}
		title := args[0]
		td.add(title)
		saveTodos(td)
		return nil
	}
	// parse args to get the title of the task to be added
	// call the add func with the title
}

func deleteCmd(td *Todos) func(args []string) error {
	return func(args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("missing id\n")
		}
		id := args[0]
		integer, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println("failed to convert string to integer", err)
		}
		td.delete(integer)
		saveTodos(td)
		return nil
	}
	// parse args to get the ID of the task to be deleted
	// call the delete func with the ID

}

func completeCmd(td *Todos) func(args []string) error {
	return func(args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("missing id\n")
		}
		done := args[0]
		integer, err := strconv.Atoi(done)
		if err != nil {
			fmt.Println("failted to convert string to integer", err)
		}
		td.complete(integer)
		saveTodos(td)
		return nil

	}
	// parse args to get the ID of the task to be completed
	// call the complete func with the ID
}

func listCmd(td *Todos) func(args []string) error {
	return func(args []string) error {
		td.showList()
		saveTodos(td)
		return nil
	}
}

func loadTodos(td *Todos) error {
	data, err := os.ReadFile("todos.json")
	if err != nil {
		if os.IsNotExist(err) {
			return nil // no file yet, start empty
		}
		return err
	}

	return json.Unmarshal(data, td)
}

func saveTodos(td *Todos) error {
	data, err := json.MarshalIndent(td, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile("todos.json", data, 0644)
}

// implementing commands

func main() {

	loadTodos(&td)

	cmdName := os.Args[1]
	args := os.Args[2:]

	for _, command := range commands {
		if command.Name == cmdName {
			err := command.Run(args)
			if err != nil {
				fmt.Println("Error:", err)
			}
			return
		}
	}
	fmt.Println("unknown command:\n", cmdName)

	saveTodos(&td)
}
