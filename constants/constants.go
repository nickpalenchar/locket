package constants

type constants struct {
	// Name of the file that will be tested with password decryption on restore (before
	// larger files are downloaded and decrypted)
	VERIFIER_FILE          string
	VERIFIER_FILE_CONTENTS string
	KEYRING                keyringConstants
}

type keyringConstants struct {
	SERVICE string
	USER    string
}

var Constants constants

func init() {
	Constants = constants{
		VERIFIER_FILE:          ".locketcheck",
		VERIFIER_FILE_CONTENTS: "OK",
		KEYRING: keyringConstants{
			SERVICE: "locket",
			USER:    "anonymous",
		},
	}
}
