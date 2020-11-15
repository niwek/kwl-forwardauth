package config

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

type env struct {
	Env            string `mapstructure:"ENV"`
	Port           string `mapstructure:"PORT"`
	AllowedOrigins string `mapstructure:"ALLOWED_ORIGINS"`
}

// Env ...
var Env env

func init() {
	viper.AutomaticEnv()
	viper.SetDefault(Constant.Env, "local")

	checkAndSetTestEnv()

	if env := viper.GetString(Constant.Env); env == "local" {
		viper.SetConfigName("local")
		viper.AddConfigPath(".")
		viper.SetConfigType("env")

		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error reading in env file: %s", err.Error())
		}
	}

	if err := viper.Unmarshal(&Env); err != nil {
		log.Fatalf(
			"Error encountered unmarshalling env variables into config struct: %s",
			err.Error(),
		)
	}
}

func checkAndSetTestEnv() {
	if strings.HasSuffix(os.Args[0], ".test") {
		viper.Set(Constant.Env, "test")
		viper.Set(Constant.AllowedOrigins, "http://foo-bar")
	}
}
