package tasks

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/rodaine/table"
)

// Blueprint for one task
// This is so they all follow a certain outline
type Task struct {
	ID        int
	Title     string
	Completed bool
}

// Slice of tasks
// Holds the list of all tasks we've made
type Tasks struct {
	items  []Task
	nextID int
}

// t is an instance of the slice
// We create a pointer receiver so that it can be edited
func (t *Tasks) Add(title string) {
	item := Task{
		ID:        t.nextID,
		Title:     title,
		Completed: false,
	}
	t.nextID++
	t.items = append(t.items, item)
}

func (t *Tasks) Delete(id int) {
	for i := 0; i < len(t.items); i++ {
		if t.items[i].ID == id {
			t.items = append(t.items[:i], t.items[i+1:]...)
			fmt.Printf("Item deleted \n")
			return
		}
	}
}

func (t *Tasks) Complete(id int) {
	for i := 0; i < len(t.items); i++ {
		if t.items[i].ID == id {
			t.items[i].Completed = true
			fmt.Printf("Item completed \n")
			return
		}
	}
}

func (t *Tasks) ShowList() {
	tbl := table.New("ID", "Name", "Completed?")

	for _, tasks := range t.items {
		tbl.AddRow(
			tasks.ID,
			tasks.Title,
			tasks.Completed,
		)
	}
	tbl.Print()
}

var t Tasks

type Command struct {
	Command string
	Help    string
	Func    func(args []string) error
}

var commands = []Command{
	{Command: "add", Help: "add a task to the list", Func: addCmd(&t)},
	{Command: "delete", Help: "delete a task from the list", Func: deleteCmd(&t)},
	{Command: "complete", Help: "mark task complete", Func: completeCmd(&t)},
	{Command: "list", Help: "shows full list of to do items", Func: showCmd(&t)},
}

func addCmd(t *Tasks) func(args []string) error {
	return func(args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("missing title\n")
		}
		title := args[0]
		t.Add(title)
		return nil
	}
}

func deleteCmd(t *Tasks) func(args []string) error {
	return func(args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("Specify ID for deletion \n")
		}
		id := args[0]
		ID, err := strconv.Atoi(id)
		if err != nil {
			fmt.Printf("failed to convert integer to string %v", err)
		}
		t.Delete(ID)
		fmt.Printf("Item deleted \n")
		return nil
	}
}

func completeCmd(t *Tasks) func(args []string) error {
	return func(args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("Specify ID to mark complete \n")
		}
		done := args[0]
		Complete, err := strconv.Atoi(done)
		if err != nil {
			fmt.Printf("failed to convert integer to string %v", err)
		}
		t.Complete(Complete)
		fmt.Printf("Item marked complete \n")
		return nil
	}
}

func showCmd(t *Tasks) func(args []string) error {
	return func(args []string) error {
		t.ShowList()
		return nil
	}
}

func main() {
	// Assign the indicies to the args
	// 0 for the command
	// 1 for the info

	// assign var called cmdName
	// if the cmdName which is of the same command indicies matches Command.Name
	// Run that command in our command directory
}

// REMEmBER: its in JSON
// Read file
func loadTasks(t *Tasks) error {
	data, err := os.ReadFile("todos.json")
	if err != nil {
		if os.IsNotExist(err) { // if there is an error a main one is to check whether it exists anymore
			return nil
		}
		return err
	}
	return json.Unmarshal(data, t)
}

// Write to file
func saveTasks() {

}
