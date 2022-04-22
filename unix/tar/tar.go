package tar

import (
	"bytes"
	"fmt"
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
		fmt.Errorf("Error occured: %v", err)
	}
	return &out
}
