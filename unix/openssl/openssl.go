/*
package openssl implements some commands from the openssl
unix command
*/
package openssl

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os/exec"
)

/*
Enc encrypts incoming data using a provided password.
Encrypted result is base64 encoded
*/
func Enc(stdin *bytes.Buffer, pw string) *bytes.Buffer {
	cmd := exec.Command("openssl", "enc", "-aes-256-cbc", "-pass", fmt.Sprintf("pass:%s", pw), "-base64")
	cmd.Stdin = stdin

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Run()

	return &out
}

/*
Dec decrypts incoming base64 data using a provided password
*/
func Dec(stdin io.Reader, pw string) *bytes.Buffer {
	cmd := exec.Command("openssl", "enc", "-d", "-aes-256-cbc", "-pass", fmt.Sprintf("pass:%s", pw), "-base64")
	cmd.Stdin = stdin

	var (
		out  bytes.Buffer
		eout bytes.Buffer
	)
	cmd.Stdout = &out
	cmd.Stderr = &eout
	err := cmd.Run()

	if err != nil {
		log.Fatalf("Error Decrypting: %s %s", eout.String(), err)
	}

	return &out
}
