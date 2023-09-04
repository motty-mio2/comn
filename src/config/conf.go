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

	_, err:=getConfigFile()

	if err != nil {

	}
}

func CreateEmptyConfig(){
	config.Set("compose_dir", "")
	config.Set("current_dir", "")

	WriteConfig()
}

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

func ReadConfig() MyConfig {
	var myconfig MyConfig

	config_file := getConfigFile(true)

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
	config_file := getConfigFile(false)

	buf := new(bytes.Buffer)

	_, err := config.DumpTo(buf, config.Toml)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	os.WriteFile(config_file, buf.Bytes(), 0755)
}
