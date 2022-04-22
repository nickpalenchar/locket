/*
Package unix contains shell commands from standard
unix builtints. They operate well with pipes and can
be chained together
*/

package unix

import (
	"bytes"
	"io/ioutil"
	"os"
)

/* ToFile writes the result of stdout to a file
of a given name. Meant to be used at the end of a
series of piped unix commands
*/
func ToFile(buffer *bytes.Buffer, filename string) {
	data, _ := ioutil.ReadAll(buffer)
	os.WriteFile(filename, data, 0666)
}
