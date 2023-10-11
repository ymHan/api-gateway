package config

import (
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	Port        string `mapstructure:"PORT"`
	UserQSvcUrl string `mapstructure:"USER_QUERY_SVC_URL"`
	UserCSvcUrl string `mapstructure:"USER_COMMAND_SVC_URL"`
	//SportsQSvcUrl string `mapstructure:"SPORTS_QUERY_SVC_URL"`
	//SportsCSvcUrl string `mapstructure:"SPORTS_COMMAND_SVC_URL"`
	//ImsQSvcUrl string `mapstructure:"IMS_QUERY_SVC_URL"`
	//ImsQSvcUrl string `mapstructure:"IMS_QUERY_SVC_URL"`
	//CmsQSvcUrl string `mapstructure:"CMS_QUERY_SVC_URL"`
	//CmsQSvcUrl string `mapstructure:"CMS_QUERY_SVC_URL"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./")

	path := ".env"

	if os.Getenv("IS_DOCKER") == "true" || os.Getenv("IS_DOCKER") == "1" {
		path = ".env.docker"
	}

	viper.SetConfigName(path)
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
