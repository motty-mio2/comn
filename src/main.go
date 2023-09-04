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
		file := cli.PickupComposeSpace(conf.ComposeDir)

		config.UpdateConfig("current_dir", file)

		config.WriteConfig()

	} else {
		args := strings.Join(os.Args[1:], " ")

		cli.ComposeWrapper(conf.ComposeDir, conf.CurrentDir, args)
	}
}
