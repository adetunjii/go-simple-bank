package util

import "github.com/spf13/viper"

type Config struct {
	DBDriver string		   		`mapstructure:"DB_DRIVER"`
	DBSource string				`mapstructure:"DB_SOURCE"`
	ServerAddress string		`mapstructure:"SERVER_ADDRESS"`
}


func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")


	/*
		Tells viper to load values from the environment variables
		and overwrite the ones loaded from the file
	*/
	viper.AutomaticEnv()


	// tells viper to start reading the config
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}