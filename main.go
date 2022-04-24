package main

import (
	"fmt"
	"locket/aws"
	"locket/cli"
	"locket/configloader"
	"locket/unix/openssl"
	"locket/unix/tar"
)

/*
main (for now) encrypts and uploads a directory tester
to an s3 bucket

tester must be under the home directory (i.e. ~/tester)

A config file named `.locket.conf.yaml` must also be
present under the home directory. See docs/.locket.conf.yaml
for an example.

For s3 upload to work correctly, a profile must be set with the
correct s3 permissions and referenced in .locket.conf.yaml.
See docs/aws-config.md
*/
func main() {
	cli := cli.NewCli()
	cli.Register("hello", hello)
	cli.Run()
}

func hello() int {
	cli.Print("hello world")
	return 0
}

func mainupload() {
	opts := configloader.Config()

	a := tar.Create("~/tester")
	encrypted := openssl.Enc(a, "tester")

	aws.UploadToS3(
		encrypted,
		opts.Auth.Aws.Bucket,
		opts.Auth.Aws.Profile,
		map[string]string{},
	)

	fmt.Println("Done 🔐")
}
