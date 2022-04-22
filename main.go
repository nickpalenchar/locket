package main

import (
	"archiver/configloader"
	"archiver/unix"
	"archiver/unix/openssl"
	"archiver/unix/tar"
	"fmt"
)

func main() {
	a := tar.Create("/Users/nick/tester")
	encrypted := openssl.Enc(a, "tester")

	fmt.Printf("enccc %s", encrypted)

	decrypted := openssl.Dec(encrypted, "tester")

	unix.ToFile(decrypted, "/Users/nick/dec.tar.gz")

	fmt.Println("hello world")
	opts := configloader.Config()
	fmt.Println("git here ", opts.Directories)
}
