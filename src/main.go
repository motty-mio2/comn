package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/motty-mio2/dockern/src/cli"
	"github.com/motty-mio2/dockern/src/config"
)

func main() {
	conf := config.ReadConfig()

	if len(os.Args) == 1 {
		fmt.Println(conf.CurrentDir)
	} else if os.Args[1] == "set" {
		files, _ := os.ReadDir(conf.ComposeDir)

		var names []string
		for _, entry := range files {
			names = append(names, entry.Name())
		}

		file := cli.NameSelector(names)

		config.WriteConfig(file)

	} else {
		args := strings.Join(os.Args[1:], " ")

		cli.ComposeWrapper(conf.ComposeDir, conf.CurrentDir, args)
	}
}
