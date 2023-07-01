package config

import "strconv"

type Mysql struct {
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	Db           string `yaml:"db"`
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	Config       string `yaml:"config"`
	LogLevel     string `yaml:"log_level"`
	MaxIdleConns int    `yaml:"max_idle_conns"`
	MaxOpenConns int    `yaml:"max_open_conns"`
	MaxLifetime  int64  `yaml:"max_lifetime"`
}

func (m *Mysql) Dsn() string {
	return m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")/" + m.Db + "?" + m.Config
}
