package types

type Command struct {
	Name        string
	Description string
	Handler     func()
}
