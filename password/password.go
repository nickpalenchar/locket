package password

import (
	"locket/cli"
	"locket/constants"

	"github.com/zalando/go-keyring"
)

type InvalidPasswordType struct{}

func (i *InvalidPasswordType) Error() string {
	return "Invalid password type provided"
}

/*
GetPassword returns the user's password based on the type.
See configloader's passwordType options for available types.
If an invalid type is provided, it returns an empty string.
*/
func GetPassword(t string) string {
	if t == "prompt" {
		return promptPassword(false)
	}

	return ""
}

func promptPassword(checkKeychain bool) string {

	if checkKeychain {
		pw := GetKeychainPassword()
		if pw != "" {
			return pw
		}
	}

	// no saved password, so ask for it
	for {
		pass1 := cli.PromptPass("Enter a password: ")
		pass2 := cli.PromptPass("Re-enter to verify: ")
		if pass1 == pass2 {
			cli.Print("Hint: you can set a default password with `locket password set` and avoid this prompt")
			return pass1
		}
		cli.Print("Passwords do not match.")
	}

}

func SetKeychainPassword(pw string) error {
	err := keyring.Set(
		constants.Constants.KEYRING.SERVICE,
		constants.Constants.KEYRING.USER,
		pw,
	)
	return err
}

/* getKeychainPassword gets the saved password on the keychain
using default config values. Returns an empty string if its not set
*/
func GetKeychainPassword() string {
	pw, err := keyring.Get(
		constants.Constants.KEYRING.SERVICE,
		constants.Constants.KEYRING.USER,
	)
	if err != nil {
		return ""
	}
	return pw
}
