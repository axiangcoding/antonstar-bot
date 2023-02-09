package settings

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

var Config ConfigStruct

func Setup() {
	setDefault()
	viper.SetConfigName("app.toml")
	viper.SetConfigType("toml")
	viper.AddConfigPath("config/")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("as")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("Config file not found.")
		}
		log.Fatal(err)
	}
	err := viper.Unmarshal(&Config)
	if err != nil {
		log.Fatalf("Config Properties unable to decode into struct, %v", err)
	}
}

func setDefault() {
	// generate a default config file maybe?
	// 是否要生成一个默认的配置文件？
}
