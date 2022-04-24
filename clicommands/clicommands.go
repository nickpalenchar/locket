/*
package clicommands contains the higher-order functions to be ran by the cli.
it imports other functions as nescessary
*/
package clicommands

import (
	"locket/cli"
	"locket/clicommands/debugcommands"
)

func AddCommands(c *cli.Cli) {
	c.Register("init", commandInit)
	
	debugcommands.AddDebugCommands(c)
}

func commandInit() int {
	cli.Print("init here")
	return 0
}
