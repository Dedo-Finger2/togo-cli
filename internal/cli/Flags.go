package cli

import (
	"flag"
	"fmt"
	"os"
	"text/tabwriter"
)

type Cli struct {
}

type Flag struct {
	Name            string
	Description     string
	VariablePointer *string
}

var Flags = []Flag{}

func (c *Cli) AddFlag(name, description string, variablePointer *string) {
	flag := Flag{
		Name:            name,
		Description:     description,
		VariablePointer: variablePointer,
	}

	Flags = append(Flags, flag)
}

func (c *Cli) Help() {
	fmt.Println("------------------------------------------------")
	fmt.Println("# TOGO - CLI")
	fmt.Println("------------------------------------------------")

	writer := tabwriter.NewWriter(os.Stdout, 0, 10, 2, ' ', tabwriter.Debug)

	defer writer.Flush()

	// helpCommands := map[string]string{
	// 	"create --name={NAME}":  "creates a new todo list.",
	// 	"add --task={NAME}":     "adds a task into your todo list.",
	// 	"delete --id={INT}":     "deletes a given task by it's id.",
	// 	"complete --id={INT}":   "completes a given task by it's id.",
	// 	"uncomplete --id={INT}": "uncompletes a given task by it's id.",
	// 	"list":                  "lists all uncompleted tasks.",
	// 	"list --all":            "lists all tasks.",
	// 	"list --completed":      "lists all completed tasks.",
	// }

	fmt.Fprint(writer, "FLAG NAME\t FLAG DESCRIPTION\n")
	for _, flag := range Flags {
		fmt.Fprint(writer, "> --", flag.Name, " \t ", flag.Description, "\n")
	}
}

func (c *Cli) Run() {
	for _, f := range Flags {
		flag.StringVar(f.VariablePointer, f.Name, "", f.Description)
	}

	flag.Parse()
}
