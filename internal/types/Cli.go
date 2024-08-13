package types

type Function = func()

type Cli interface {
	AddFlag(name, description string, variablePointer *string)
	AddCommand(name, description string, handler Function)
	Help()
	ParseFlags()
	Start()
}
