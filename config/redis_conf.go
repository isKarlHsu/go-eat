package config

import "strconv"

type Redis struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
	PoolSize int    `yaml:"pool_size"`
}

func (r Redis) Addr() string  {
	return r.Host + ":" + strconv.Itoa(r.Port)
}