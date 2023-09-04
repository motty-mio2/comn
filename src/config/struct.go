package config

type MyConfig struct {
	Backend    string `config:"backend"`
	ComposeDir string `config:"compose_dir"`
	CurrentDir string `config:"current_dir"`
}
