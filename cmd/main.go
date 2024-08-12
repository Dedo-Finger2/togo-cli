package main

import (
	"flag"

	cliPkg "github.com/Dedo-Finger2/todo-list-cli/internal/cli"
	"github.com/Dedo-Finger2/todo-list-cli/internal/commands"
)

var cli = cliPkg.Cli{}

func init() {
	var (
		toDoListName       string
		taskName           string
		taskID             string
		listCompletedTasks string
	)

	flag.Usage = cli.Help

	cli.AddFlag("name", "sets a name for your todo list.", &toDoListName)
	cli.AddFlag("task", "sets a name for a task.", &taskName)
	cli.AddFlag("id", "chooses a task id", &taskID)
	cli.AddFlag("completed", "list only completed tasks", &listCompletedTasks)

	cli.AddCommand("create", "Creates a new to go task", commands.CreateToGoList)
	cli.AddCommand("add", "Adds a new task", commands.AddTask)
	cli.AddCommand("list", "Lists the tasks", commands.ListTasks)
	cli.AddCommand("complete", "Completes a task", commands.CompleteTask)
	cli.AddCommand("incomplete", "Incompletes a task", commands.IncompleteTask)
	cli.AddCommand("delete", "Deletes a task", commands.DeleteTask)

	cli.ParseFlags()
}

func main() {
	cli.Start()
}
