package config

import "github.com/spf13/viper"

var Manager *Config

type Config struct {
	ApiKey            string `mapstructure:"API_KEY"`
	DictionaryBaseUrl string `mapstructure:"DICTIONARY_BASE_URL"`
}

func LoadConfig(path string) (err error) {
	Manager = &Config{}
	viper.AddConfigPath(path)
	viper.SetConfigName("app.env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(Manager)
	return
}
