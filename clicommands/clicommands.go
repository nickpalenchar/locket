/*
package clicommands contains the higher-order functions to be ran by the cli.
it imports other functions as necessary
*/
package clicommands

import (
	"fmt"
	"locket/cli"
	"locket/clicommands/debugcommands"
	"locket/configloader"
	"os"
	"path"
)

func AddCommands(c *cli.Cli) {
	c.Register("init", commandInit, &cli.CliOpts{
		Help: "Interactively generate config file",
	})

	c.Register("backup", commandBackup, &cli.CliOpts{
		Help: "Create an encrypted backup to s3",
	})

	c.Register("restore", commandRestore, &cli.CliOpts{
		Help: "Restore files from a backup",
	})

	debugcommands.AddDebugCommands(c)
}

func commandInit() int {
	var (
		directory string
		profile   string
		bucket    string
	)

	if _, err := os.Stat(configloader.ConfigPath()); err == nil {
		cli.Print("Warning! There is already a config file present, continuing will overwrite.")
		cli.Print("Type \"yes\" to continue...")
		answer := cli.Prompt("> ")
		if answer != "yes" {
			cli.Print("Didn't say yes. Canceling")
			return 0
		}
	}

	cli.Print("Directory to backup (e.g. ~)")
	cli.Print("(you can add more later)")
	directory = cli.Prompt("> ")

	cli.Print("Aws bucket to use:")
	bucket = cli.Prompt("> ")

	cli.Print("Aws profile to auth with bucket (e.g. default)")
	profile = cli.Prompt("> ")

	configopts := configloader.NewConfig(directory, profile, bucket)
	fmt.Printf("%v", configopts)
	configopts.ToFile(path.Join(os.Getenv("HOME"), ".locket.conf.yaml"))

	cli.Print(fmt.Sprintf("Config written to %s", configloader.ConfigPath()))

	return 0
}
