package configs

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Port              string `envconfig:"PORT"`
	GoogleAppCertPath string `envconfig:"GOOGLE_CERT_PATH"`
	GoogleAppCertJSON string `envconfig:"GOOGLE_CERT"`
}

var Global Config

func init() {
	var config Config
	if err := envconfig.Process("", &config); err != nil {
		panic(err)
	}
	Global = config
}
