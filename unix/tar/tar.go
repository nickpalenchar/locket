package tar

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
)

/*
Create makes a .tar.gz archive of the provided directory
*/
func Create(path string) *bytes.Buffer {
	basepath := path
	cmd := exec.Command("tar", "-czvf", "-", basepath)
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		fmt.Errorf("Error occurred: %v", err)
	}
	return &out
}

/* Extract runs `tar -xzf` on the provided io and extracts
it relative to the given path (p). If p is ".", it extracts
relative to the current working directory. Directories are created
as needed if non-existant ones are supplied to p.
*/
func Extract(data io.Reader, p string) error {
	var cmd *exec.Cmd
	if p == "." {
		cmd = exec.Command("tar", "-xf", "-")
	} else {
		if err := exec.Command("mkdir", "-p", p).Run(); err != nil {
			return err
		}
		cmd = exec.Command("tar", "-xf", "-", "-C", p)
	}
	cmd.Stdin = data
	return cmd.Run()
}
