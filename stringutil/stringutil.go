package stringutil

import "strings"

func PathPos(path string, n int) string {
	return strings.Split(path, "/")[n]
}
