package config

type Email struct {
	Host             string `json:"host" yaml:"host"`
	Port             int    `json:"port" yaml:"port"`
	User             string `json:"user" yaml:"user"`
	Password         string `json:"password" yaml:"password"`
	DefaultFromEmail string `json:"default_from_email" yaml:"default_from_email"`
	UseSSL           bool   `json:"use_ssl" yaml:"use_ssl"`
	UserTls          bool   `json:"user_tls" yaml:"user_tls"`
}
