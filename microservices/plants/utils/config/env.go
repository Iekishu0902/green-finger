package config

import "github.com/spf13/viper"

func init() {
	viper.AutomaticEnv()
}

func GetString(key string) string {
	return viper.GetString(key)
}
