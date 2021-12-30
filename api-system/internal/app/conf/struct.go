package conf

type AllConfig struct {
	App
	Server
}

type App struct {
	Version string
	Name    string
	Log     struct {
		Level string
		File  struct {
			Enable bool
			Path   string
		}
	}
	Token struct {
		Secret          string
		ExpireDuration  string `mapstructure:"expire_duration"`
		RefreshDuration string `mapstructure:"refresh_duration"`
	}
	Swagger struct {
		Enable bool
	}
	Data struct {
		Database struct {
			Driver      string
			Source      string
			MaxIdleConn int `mapstructure:"max_idle_conn"`
			MaxOpenConn int `mapstructure:"max_open_conn"`
		}
		Cache struct {
			Driver string
			Source string
		}
	}
	Response struct {
		HideErrorDetails bool `mapstructure:"hide_error_details"`
	}
}

type Server struct {
	RunMode  string `mapstructure:"run_mode"`
	Port     string
	BasePath string `mapstructure:"base_path"`
}
