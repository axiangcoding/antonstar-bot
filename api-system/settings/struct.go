package settings

type ConfigStruct struct {
	App
	Server
}

type App struct {
	Version string
	Name    string
	Log     struct {
		Level string
		File  struct {
			Enable  bool
			Path    string
			Encoder string
		}
	}
	Auth struct {
		Secret              string
		ExpireDuration      string `mapstructure:"expire_duration"`
		RefreshDuration     string `mapstructure:"refresh_duration"`
		CasbinModelPath     string `mapstructure:"casbin_model_path"`
		CasbinPolicyAdapter string `mapstructure:"casbin_policy_adapter"`
		CasbinPolicyPath    string `mapstructure:"casbin_policy_path"`
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
	MQ struct {
		Source string
	}
	Response struct {
		HideErrorDetails bool `mapstructure:"hide_error_details"`
	}
	Upload struct {
		SuperBed struct {
			Token      string
			Categories string
		}
	}
	CqHttp struct {
		Url    string `mapstructure:"url"`
		SelfQQ int64  `mapstructure:"self_qq"`
	}
}

type Server struct {
	RunMode  string `mapstructure:"run_mode"`
	Port     string
	BasePath string `mapstructure:"base_path"`
}
