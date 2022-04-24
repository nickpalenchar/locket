/*
package cli contains utilities for parsing command line arguments
and associated flags, as well as functions to register other
functions as associated with the cli
*/
package cli

import (
	"fmt"
	"os"
)

type CliCommand struct {
	Name string
	Func CliFunction
}

/*
CliFunction is a function that return
an int representing the exit code in a
shell environment
*/
type CliFunction func() int

type Cli struct {
	commands map[string]CliCommand
}

func NewCli() *Cli {
	commands := make(map[string]CliCommand)
	return &Cli{
		commands: commands,
	}
}

/* Register registers a function on the cli, so it
can be invoked by name via command line arguments */
func (c *Cli) Register(name string, f CliFunction) {

	_, exists := c.commands[name]

	if exists {
		panic(fmt.Sprintf("cannot register function name %s when it already exists", name))
	}

	c.commands[name] = CliCommand{
		Name: name,
		Func: f,
	}
}

func (c *Cli) Run() (code int) {
	fmt.Println("time to run stuff")
	args := os.Args[1:]

	if len(args) == 0 {
		cliHelp()
		return 1
	}

	cmd, ok := c.commands[args[0]]

	if !ok {
		cliHelp()
		return 1
	}

	code = cmd.Func()

	return code
}

/*
Print takes a string and prints it to the terminal
*/
func Print(s string) {
	fmt.Println(s)
}

func cliHelp() {
	fmt.Println("Invalid command")
	// TODO list commands available
}
