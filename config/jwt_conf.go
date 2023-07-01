package config

type Jwt struct {
	Secret string `yaml:"secret"`
	Expires int `yaml:"expires"`
	Issuer string `yaml:"issuer"`
}
