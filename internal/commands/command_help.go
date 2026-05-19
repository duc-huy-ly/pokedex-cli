package commands

import (
	"fmt"
)

type CommandHelp struct {
	CliCommand
}

func NewCommandHelp() *CommandHelp {
	return &CommandHelp{
		CliCommand: CliCommand{
			Name: "help",
			Description: "Displays the help",
		},
	}
}
func (c *CommandHelp) Execute() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, cmd := range GetCommands(nil) {
		cmd.DisplayInfo()
	}
	fmt.Println()
	return nil
}

