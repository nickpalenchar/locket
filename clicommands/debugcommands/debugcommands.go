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

var s3Client aws.S3Client
var conf configloader.Configopts

func init() {
	conf = *configloader.Config()
	s3Client = aws.NewS3Client(conf.Auth.Aws.Profile, conf.Auth.Aws.Bucket)
}

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

	a := tar.Create("~/tester")
	encrypted := openssl.Enc(a, "tester")

	s3Client.Upload(
		encrypted,
		"tester",
		map[string]string{},
	)

	cli.Print("Done 🔐")
	return 0
}
