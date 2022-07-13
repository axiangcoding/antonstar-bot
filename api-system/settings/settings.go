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
	viper.AddConfigPath("config/default/")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("ax")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("Config file not found. The program will use default conf and may not work properly")
		}
	}
	err := viper.Unmarshal(&Config)
	if err != nil {
		log.Fatalf("Config Properties unable to decode into struct, %v", err)
	}
	// marshal, _ := json.MarshalIndent(Config, "-- ", "  ")
	// log.Println("Effective configuration: " + string(marshal))
}

func setDefault() {
	// generate a default config file maybe?
	// 是否要生成一个默认的配置文件？
}
