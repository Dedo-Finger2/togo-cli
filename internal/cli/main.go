package cli

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	types "github.com/Dedo-Finger2/todo-list-cli/internal/types"
)

type Cli struct{}

var (
	Flags    = []types.Flag{}
	Commands = []types.Command{}
)

func (c *Cli) AddFlag(name, description string, variablePointer *string) {
	flag := types.Flag{
		Name:            name,
		Description:     description,
		VariablePointer: variablePointer,
	}

	Flags = append(Flags, flag)
}

func (c *Cli) AddCommand(name, description string, handler func()) {
	command := types.Command{
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

	fmt.Fprint(writer, "FLAG NAME\t FLAG DESCRIPTION\n")
	for _, flag := range Flags {
		fmt.Fprint(writer, "> --", flag.Name, " \t ", flag.Description, "\n")
	}

	fmt.Fprint(writer, "\nCOMMAND NAME\t COMMAND DESCRIPTION\n")
	for _, command := range Commands {
		fmt.Fprint(writer, "> ", command.Name, " \t ", command.Description, "\n")
	}
}

func (c *Cli) ParseFlags() {
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
