package configs

import "context"

// HOME OF ALL GLOBAL VARIABLES
type Config struct {
	USER     string `mapstructure:"USER"`
	PASSWORD string `mapstructure:"PASSWORD"`
	IP       string `mapstructure:"IP"`
	DB       string `mapstructure:"DB"`
	PORT     string `mapstructure:"PORT"`
}

var CONFIG Config

var DBCtx context.Context
