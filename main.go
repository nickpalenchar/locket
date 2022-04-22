package main

import (
	"archiver/configloader"
	"archiver/unix/tar"
	"fmt"
)

func main() {
	tar.Create("/Users/nick/tester")
	fmt.Println("hello world")
	opts := configloader.Config()
	fmt.Println("git here ", opts.Directories)
}
