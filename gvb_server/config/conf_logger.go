package config

type Logger struct {
	Level        string `yaml:"level"`
	Prefix       string `yaml:"prefix"`
	Director     string `yaml:"director"`
	ShowLine     bool   `yaml:"show_Line"`
	LogInConsole bool   `yaml:"log_In_Console"`
}
