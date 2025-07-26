package config

import (
	"github.com/spf13/viper"
	"log"
)

/*
ParseConfig Config parser function
  - [conf] : the structure of your config file
  - [type_of] : the type of config file (yaml, env, json etc)
  - [file_path] : the path of your config file
  - [file_name] : the file name of your config file
  - you can find the example config.yaml and models.go file under config folder
*/
func ParseConfig(conf interface{}, type_of, file_path, file_name string) (err error) {
	viper.SetConfigType(type_of)
	viper.AddConfigPath(file_path)
	configPath := file_name

	viper.SetConfigName(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	err = viper.Unmarshal(&conf)
	if err != nil {
		log.Printf("Unable to decode config file into struct, %v \n", err)
		return err
	}

	return nil
}
