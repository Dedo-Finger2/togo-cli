package cli

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

type Cli struct {
}

type Flag struct {
	Name            string
	Description     string
	VariablePointer *string
}

type Command struct {
	Name        string
	Description string
	Handler     func()
}

var (
	Flags    = []Flag{}
	Commands = []Command{}
)

func (c *Cli) AddFlag(name, description string, variablePointer *string) {
	flag := Flag{
		Name:            name,
		Description:     description,
		VariablePointer: variablePointer,
	}

	Flags = append(Flags, flag)
}

func (c *Cli) AddCommand(name, description string, handler func()) {
	command := Command{
		Name:        name,
		Description: description,
		Handler:     handler,
	}

	Commands = append(Commands, command)
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

	fmt.Fprint(writer, "\nCOMMAND NAME\t COMMAND DESCRIPTION\n")
	for _, command := range Commands {
		fmt.Fprint(writer, "> ", command.Name, " \t ", command.Description, "\n")
	}
}

func (c *Cli) ParseFlagsAndCommands() {
	for _, f := range Flags {
		flag.StringVar(f.VariablePointer, f.Name, "", f.Description)
	}

	flag.Parse()
}

func (c *Cli) Start() {
	input := strings.ToUpper(flag.Arg(0))
	commandFound := false

	if len(flag.Args()) > 0 {
		for _, command := range Commands {
			if input == strings.ToUpper(command.Name) {
				command.Handler()
				commandFound = true
				break
			}
		}

		if !commandFound {
			fmt.Println("Command not found.")
		}
	} else {
		fmt.Println("Use the flag '--help' to see the avaliables commands.")
	}
}
