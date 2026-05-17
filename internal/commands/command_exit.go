package commands

import (
	"fmt"
	"os"
)

func CommandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return fmt.Errorf("os.Exit didn't close properly")
}