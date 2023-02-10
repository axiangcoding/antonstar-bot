package settings

import "github.com/gin-gonic/gin"

const (
	AppRunModeRelease = gin.ReleaseMode
	AppRunModeDebug   = gin.DebugMode
)

const (
	AppLogFileEncoderJson    = "json"
	AppLogFileEncoderConsole = "console"
)

type GlobalConf struct {
	App    appConf
	Server serverConf
}

type serverConf struct {
	RunMode  string `mapstructure:"run_mode"`
	Port     string `mapstructure:"port"`
	BasePath string `mapstructure:"base_path"`
}

type appConf struct {
	Version string `mapstructure:"version"`
	Name    string `mapstructure:"name"`
	Log     struct {
		Level string
		File  struct {
			Dir     string `mapstructure:"dir"`
			Encoder string `mapstructure:"encoder"`
		}
	}
	Auth struct {
		Session struct {
			EncryptSecret string `mapstructure:"encrypt_secret"`
			MaxAge        string `mapstructure:"max_age"`
		}
	}
	Swagger struct {
		Enable bool `mapstructure:"enable"`
	}
	Data struct {
		Db struct {
			Source      string `mapstructure:"source"`
			MaxIdleConn int    `mapstructure:"max_idle_conn"`
			MaxOpenConn int    `mapstructure:"max_open_conn"`
		}
		Cache struct {
			Source string `mapstructure:"source"`
		}
	}
	Service struct {
		CqHttp struct {
			Url    string `mapstructure:"url"`
			SelfQQ int64  `mapstructure:"self_qq"`
			Secret string `mapstructure:"secret"`
		}
	}
}
