package commands

import (
	"fmt"
)


type ICommand interface {
	Execute() error
	DisplayInfo()
}
type CliCommand struct {
	Name        string
	Description string
	Arguments   []string
}

func (c *CliCommand) DisplayInfo (){
	fmt.Printf("%s : %s\n", c.Name, c.Description)
}
