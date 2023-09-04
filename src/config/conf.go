package config

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/toml"
)

func InitConfig() {
	config.WithOptions(config.ParseEnv)
	config.WithOptions(func(opt *config.Options) {
		opt.DecoderConfig.TagName = "config"
	})

	config.AddDriver(toml.Driver)
}

func getConfigFile() string {
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
		fmt.Println("config file does not exist")
	}

	return config_file
}

func ReadConfig() MyConfig {
	var myconfig MyConfig

	config_file := getConfigFile()

	err := config.LoadFiles(config_file)
	if err != nil {
		panic(err)
	}

	err = config.BindStruct("", &myconfig)
	if err != nil {
		panic(err)
	}

	return myconfig
}

func UpdateConfig(key string, value string) {
	config.Set(key, value)

}

func WriteConfig() {
	config_file := getConfigFile()

	buf := new(bytes.Buffer)


	_, err := config.DumpTo(buf, config.Toml)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	os.WriteFile(config_file, buf.Bytes(), 0755)
}
