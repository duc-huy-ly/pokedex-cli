package commands

import (
	"fmt"
	"os"
)

type ExitCommand struct {
	CliCommand
}


func (c *ExitCommand) Execute() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return fmt.Errorf("ExitCommand did not close properly\n")
}

