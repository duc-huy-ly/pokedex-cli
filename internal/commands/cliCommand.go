package commands

import "fmt"


type ICommand interface {
	Execute() error
}
type CliCommand struct {
	Name        string
	Description string
	Arguments   []string
}

func (c *CliCommand) Execute() error{
	fmt.Println("Called execute from parent")
	return nil
}