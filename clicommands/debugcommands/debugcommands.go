/*
package debugcommands contains additional commands
not intended to be show to the user, but useful for
testing/debugging in development.

They become available when the environment variable
__LOCKET_DEBUG is set.
*/

package debugcommands

import (
	"fmt"
	"locket/aws"
	"locket/cli"
	"locket/configloader"
	"locket/unix/openssl"
	"locket/unix/tar"
	"os"
	"strings"
)

func AddDebugCommands(c *cli.Cli) {
	if os.Getenv("__LOCKET_DEBUG") != "" {
		c.Register("d--hello", commandHello, &cli.CliOpts{})
		c.Register("d--test-upload", testUpload, &cli.CliOpts{})
	}
}

func commandHello() int {
	cli.Print("Hello world!")
	s := strings.ReplaceAll("hello/world:hh:ss.hhh", "/:", "_")
	fmt.Println(s)
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
		"tester",
		map[string]string{},
	)

	cli.Print("Done 🔐")
	return 0
}
