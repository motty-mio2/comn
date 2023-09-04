package functions

import (
	"os/user"
	"strings"
)

func ExpandHome (path string) string {

	usr, _ := user.Current()

	return strings.Replace(path, "~", usr.HomeDir, 1)
}