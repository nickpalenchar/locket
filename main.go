package main

import (
	"archiver/aws"
	"archiver/configloader"
	"archiver/unix/openssl"
	"archiver/unix/tar"
	"fmt"
)

func main() {
	opts := configloader.Config()

	a := tar.Create("/Users/nick/tester")
	encrypted := openssl.Enc(a, "tester")

	aws.UploadToS3(
		encrypted,
		opts.Auth.Aws.Bucket,
		opts.Auth.Aws.Profile,
		map[string]string{},
	)

	fmt.Println("hello world")
	fmt.Println("git here ", opts.Directories)
}
