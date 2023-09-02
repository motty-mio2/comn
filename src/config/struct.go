package config

type MyConfig struct {
	ComposeDir string `config:"compose_dir"`
	CurrentDir string `config:"current_dir"`
}
