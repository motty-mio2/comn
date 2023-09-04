package config

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/toml"
)

func getConfigFile() string {
	conf, err := os.UserConfigDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	config_dir := filepath.Join(conf, "comn")

	if _, err := os.Stat(config_dir); os.IsNotExist(err) {
		os.Mkdir(config_dir ,0700)
	}

	config_file := filepath.Join(config_dir, "config.toml")

	if _, err := os.Stat(config_dir); os.IsNotExist(err) {
		fmt.Println("config file does not exist")
	}

	return config_file
}

func ReadConfig() MyConfig {
	var myconfig MyConfig


	config_file := getConfigFile()

	config.WithOptions(config.ParseEnv)
	config.WithOptions(func(opt *config.Options) {
		opt.DecoderConfig.TagName = "config"
	})

	config.AddDriver(toml.Driver)

	err := config.LoadFiles(config_file)
	if err != nil {
		panic(err)
	}

	err = config.BindStruct("", &myconfig)
	if err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)

	config.Set("current_dir", "test")

	_, err = config.DumpTo(buf, config.Toml)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	os.WriteFile("my-config.toml", buf.Bytes(), 0755)

	return myconfig
}

func WriteConfig(newdir string) {

	config_file := getConfigFile()

	// オプションの設定は不要

	config.AddDriver(toml.Driver)

	buf := new(bytes.Buffer)

	config.Set("current_dir", newdir)

	_, err := config.DumpTo(buf, config.Toml)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	os.WriteFile(config_file, buf.Bytes(), 0755)
}
