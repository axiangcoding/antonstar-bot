package setting

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

var _conf *GlobalConf

func InitConf() {
	viper.SetConfigName("app.toml")
	viper.SetConfigType("toml")
	viper.AddConfigPath("config/")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("as")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}

	if err := viper.Unmarshal(&_conf); err != nil {
		log.Fatalln(err)
	}
}

func C() *GlobalConf {
	return _conf
}
