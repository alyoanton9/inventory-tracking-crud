package main

import(
	"log"
	"github.com/spf13/viper"
)

type Config struct {
	Port string `mapstructure:"port"`
}

var AppConfig *Config

func LoadAppConfig(){
	log.Println("Loading server configs")

	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	loadConfigError := viper.ReadInConfig()
	if loadConfigError != nil {
		log.Fatal(loadConfigError)
	}

	unmarshallError := viper.Unmarshal(&AppConfig)
	if unmarshallError != nil {
		log.Fatal(unmarshallError)
	}
}
