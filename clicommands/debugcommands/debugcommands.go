/*
package debugcommands contains additional commands
not intended to be show to the user, but useful for
testing/debugging in development.

They become available when the environment variable
__LOCKET_DEBUG is set.
*/

package debugcommands

import (
	"locket/aws"
	"locket/cli"
	"locket/configloader"
	"locket/unix/openssl"
	"locket/unix/tar"
	"os"
)

func AddDebugCommands(c *cli.Cli) {
	if os.Getenv("__LOCKET_DEBUG") != "" {
		c.Register("d--hello", commandHello)
		c.Register("d--test-upload", testUpload)
	}
}

func commandHello() int {
	cli.Print("Hello world!")
	return 0
}

func testUpload() int {
	opts := configloader.Config()

	a := tar.Create("~/tester")
	encrypted := openssl.Enc(a, "tester")

	aws.UploadToS3(
		encrypted,
		opts.Auth.Aws.Bucket,
		opts.Auth.Aws.Profile,
		map[string]string{},
	)

	cli.Print("Done 🔐")
	return 0
}
