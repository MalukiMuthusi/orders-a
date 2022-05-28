package main

import (
	"github.com/MalukiMuthusi/orders-a/internal/logger"
	"github.com/MalukiMuthusi/orders-a/internal/utils"
	"github.com/spf13/viper"
)

func main() {

}

func init() {
	viper.AutomaticEnv()

	viper.SetEnvPrefix(utils.AppName)

	BindEnvs()

	viper.SetDefault(utils.Port, "8080")

	CheckMustBeSetEnvs()
}

// BindEnvs binds an env to a string value
func BindEnvs() {

	viper.BindEnv(utils.Port, "PORT")

	// path to a csv file with mappings for a country name and its phone number REGEX
	viper.BindEnv(utils.CountryCodesPath)

	// path to a csv file with orders data
	viper.BindEnv(utils.OrdersPath)

}

func EnvMustBeSet(key string) {
	if !viper.IsSet(key) {
		logger.Log.WithField(key, false).Fatal("not set")
	}
}

// CheckMustBeSetEnvs checks that the provided envs have been set. The application will not start if the envs provided here have not been set
func CheckMustBeSetEnvs() {

	EnvMustBeSet(utils.CountryCodesPath)
	EnvMustBeSet(utils.OrdersPath)
}
