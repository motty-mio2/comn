package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/motty-mio2/comn/src/cli"
	"github.com/motty-mio2/comn/src/config"
)

func main() {
	config.InitConfig()

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

		config.UpdateConfig("current_dir", file)

		config.WriteConfig()

	} else {
		args := strings.Join(os.Args[1:], " ")

		cli.ComposeWrapper(conf.ComposeDir, conf.CurrentDir, args)
	}
}
