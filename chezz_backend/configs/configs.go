package configs

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func init() {
	var err error
	CONFIG, err = LoadConfig("./")
	if err != nil {
		log.Fatalf("Unable to load Environment Variables - %v", err.Error())
	}
	fmt.Println("ENV LOADED!!")
}

func LoadConfig(path string) (config Config, err error) {

	// Set default values
	// viper.SetDefault(config.PORT, "8080")

	viper.AddConfigPath(path)
	viper.SetConfigName("chezz")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error in ReadInConfig : %v\n", err.Error())
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Printf("Error in ReadInConfig : %v\n", err.Error())
		return
	}

	return
}
