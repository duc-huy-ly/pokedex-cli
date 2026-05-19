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
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range GetCommands() {
		cmd.DisplayInfo()
	}
	fmt.Println()
	return nil
}

func GetCommands() map[string]ICommand {
	return map[string]ICommand{
		"exit": NewExitCommand(),
		"help": NewCommandHelp(),	
	}
}
