package clicommands

import (
	"locket/cli"
	"locket/password"
)

func commandSetPassword() int {
	pw := password.PromptPassword(false)
	password.SetKeychainPassword(pw)
	cli.Print("Hint: if you don't want to use this password on a given backup, run `locket backup --prompt-password`")
	return 0
}
