package main

import (
	"fmt"
	"locket/aws"
	"locket/configloader"
	"locket/unix/openssl"
	"locket/unix/tar"
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
