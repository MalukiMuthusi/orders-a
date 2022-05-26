package main

import (
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

}

// CheckMustBeSetEnvs checks that the provided envs have been set. The application will not start if the envs provided here have not been set
func CheckMustBeSetEnvs() {

}
