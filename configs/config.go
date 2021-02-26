package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"sync"
)

type AppConfig struct {
	RunMode      string `mapstructure:"run_mode"`
	AppPath      string `mapstructure:"app_path"`
	TemplatePath string `mapstructure:"template_path"`
}

var App = &AppConfig{}

var once sync.Once

func init() {
	once.Do(func() {
		viper.AddConfigPath("../../configs")
		viper.SetConfigType("yaml")
		// load app.yaml
		viper.SetConfigName("app")
		err := viper.ReadInConfig()
		if err != nil {
			panic(fmt.Sprintf("load app.yaml err: %v\n", err))
		}
		viper.UnmarshalKey("app", App)
	})
}
