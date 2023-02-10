package settings

// TODO

type GlobalConf struct {
	App    appConf
	Server serverConf
}

type appConf struct {
	Version string
	Name    string
	Log     struct {
		Level string
		File  struct {
			Dir     string
			Encoder string
		}
	}
	Auth struct {
		Session struct {
			EncryptSecret string
		}
		Cookie struct {
			MaxAge string
		}
	}
	Swagger struct {
		Enable bool
	}
	Data struct {
		Database struct {
			Driver      string
			Source      string
			MaxIdleConn int
			MaxOpenConn int
		}
		Cache struct {
			Driver string
			Source string
		}
	}
	Service struct {
		CqHttp struct {
			Url    string
			SelfQQ int64
			Secret string
		}
	}
	Response struct {
		HideErrorDetails bool
	}
}

type serverConf struct {
	RunMode  string `mapstructure:"run_mode"`
	Port     string `mapstructure:"port"`
	BasePath string `mapstructure:"base_path"`
}
