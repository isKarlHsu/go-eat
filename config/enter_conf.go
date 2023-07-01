package config

type Config struct {
	System System `yaml:"system"`
	Mysql  Mysql  `yaml:"mysql"`
	Redis  Redis  `yaml:"redis"`
	Logger Logger `yaml:"logger"`
	Jwt    Jwt    `yaml:"jwt"`
}
