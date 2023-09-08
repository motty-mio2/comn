package config

import (
	"bytes"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/toml"
	"github.com/motty-mio2/comn/src/cli"
	"github.com/motty-mio2/comn/src/functions"
)


func getConfigFile() (string, error) {
	conf, err := os.UserConfigDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	config_dir := filepath.Join(conf, "comn")

	if _, err := os.Stat(config_dir); os.IsNotExist(err) {
		os.Mkdir(config_dir, 0700)
	}

	config_file := filepath.Join(config_dir, "config.toml")

	if _, err := os.Stat(config_file); os.IsNotExist(err) {
		return config_file, fmt.Errorf("file: %s does not exist", config_file)
	}

	return config_file, nil
}

func validateHomeDir(path string) string {
	usr, _ := user.Current()

	if strings.Contains(path, "~") {
		return strings.Replace(path, "~", usr.HomeDir, 1)
	}	else if !strings.HasPrefix(path, "/") {
		return filepath.Join(usr.HomeDir, path)
	}	else {
		return path
	}
}


func InitConfig() {
	config.WithOptions(config.ParseEnv)
	config.WithOptions(func(opt *config.Options) {
		opt.DecoderConfig.TagName = "config"
	})

	config.AddDriver(toml.Driver)

	_, err := getConfigFile()

	if err != nil {
		CreateEmptyConfig()
	}
}

func CreateEmptyConfig() {
	var dir string

	fmt.Print("Input compose folder : ")
	_, err := fmt.Scanf("%s", &dir)
	if err != nil {
		fmt.Println("Input Error : ", err)
		os.Exit(1)
	}

	bin := functions.FzfSelector([]string{"docker", "nerdctl"})

	config.Set("backend", bin)
	config.Set("compose_dir", dir)

	file := cli.PickupComposeSpace(dir)

	UpdateConfig("current_dir", file)

	WriteConfig()

	os.Exit(0)
}


func ReadConfig() MyConfig {
	var myconfig MyConfig

	config_file, _ := getConfigFile()

	err := config.LoadFiles(config_file)
	if err != nil {
		panic(err)
	}

	err = config.BindStruct("", &myconfig)
	if err != nil {
		panic(err)
	}

	myconfig.ComposeDir = validateHomeDir(myconfig.ComposeDir)

	return myconfig
}

func UpdateConfig(key string, value string) {
	config.Set(key, value)

}

func WriteConfig() {
	config_file, _ := getConfigFile()

	buf := new(bytes.Buffer)

	_, err := config.DumpTo(buf, config.Toml)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	os.WriteFile(config_file, buf.Bytes(), 0700)
}

