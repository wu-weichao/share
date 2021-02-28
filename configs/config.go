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

type DatabaseConfig struct {
	Type        string `mapstructure:"type"`
	Host        string `mapstructure:"host"`
	Port        int    `mapstructure:"port"`
	User        string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	DBName      string `mapstructure:"dbname"`
	TablePrefix string `mapstructure:"table_prefix"`
	Charset     string `mapstructure:"charset"`
}

var Database = &DatabaseConfig{}

type JwtConfig struct {
	Secret string `mapstructure:"secret"`
	Ttl    int    `mapstructure:"ttl"`
}

var Jwt = &JwtConfig{}

var once sync.Once

func init() {
	once.Do(func() {
		viper.AddConfigPath("../../configs")
		viper.SetConfigType("yaml")
		// load app.yaml
		viper.SetConfigName("app")
		var err error
		err = viper.ReadInConfig()
		if err != nil {
			panic(fmt.Sprintf("load app.yaml err: %v\n", err))
		}
		viper.UnmarshalKey("app", App)
		viper.UnmarshalKey("jwt", Jwt)
		// load database.yaml
		viper.SetConfigName("database")
		err = viper.ReadInConfig()
		if err != nil {
			panic(fmt.Sprintf("load database.yaml err: %v\n", err))
		}
		viper.UnmarshalKey("database", Database)
	})
}
