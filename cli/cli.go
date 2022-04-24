/*
package cli contains utilities for parsing command line arguments
and associated flags, as well as functions to register other
functions as associated with the cli
*/
package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type CliCommand struct {
	Name string
	Help string
	Func CliFunction
}

/*
CliFunction is a function that return
an int representing the exit code in a
shell environment
*/
type CliFunction func() int

type Cli struct {
	Name     string
	commands map[string]CliCommand
}

func NewCli(name string) *Cli {
	commands := make(map[string]CliCommand)
	return &Cli{
		Name:     name,
		commands: commands,
	}
}

type CliOpts struct {
	Help string
}

/* Register registers a function on the cli, so it
can be invoked by name via command line arguments */
func (c *Cli) Register(name string, f CliFunction, o *CliOpts) {

	_, exists := c.commands[name]

	if exists {
		panic(fmt.Sprintf("cannot register function name %s when it already exists", name))
	}

	c.commands[name] = CliCommand{
		Name: name,
		Func: f,
		Help: o.Help,
	}
}

func (c *Cli) Run() (code int) {
	args := os.Args[1:]

	if len(args) == 0 {
		c.Help()
		return 1
	}

	cmd, ok := c.commands[args[0]]

	if !ok {
		c.Help()
		return 1
	}

	code = cmd.Func()

	return code
}

func (c *Cli) Help() {
	fmt.Printf("Usage:\n\n")
	fmt.Printf("    %s <command>\n\n", c.Name)

	Print("Commands:\n")
	for command := range c.commands {
		fmt.Printf("    %-14s %s\n",
			command,
			c.commands[command].Help,
		)
	}
}

/*
Print takes a string and prints it to the terminal
*/
func Print(s string) {
	fmt.Println(s)
}

/*
Prompt prompts the users for a value and retruns
it
*/
func Prompt(message string) string {
	fmt.Print(message)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.Trim(input, "\n")
}
