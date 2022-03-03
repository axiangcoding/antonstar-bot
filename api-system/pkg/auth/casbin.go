package auth

import (
	"axiangcoding/antonstar/api-system/internal/app/conf"
	"axiangcoding/antonstar/api-system/internal/app/data"
	"axiangcoding/antonstar/api-system/pkg/logging"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"path"
)

var enforcer *casbin.Enforcer

func SetupCasbin() {

	modelPath := conf.Config.Auth.CasbinModelPath
	policyPath := conf.Config.Auth.CasbinPolicyPath
	adapter := conf.Config.Auth.CasbinPolicyAdapter
	var e *casbin.Enforcer

	switch adapter {
	case "db":
		// 使用 gormadapter 来从数据库中加载policy，并且使用的是已经存在的gorm实例
		adapter, err := gormadapter.NewAdapterByDB(data.GetDB())
		if err != nil {
			logging.Fatal(err)
		}
		e, err = casbin.NewEnforcer("config/default/model.conf", adapter)
		if err != nil {
			logging.Fatal(err)
		}
		break
	case "local":
		var err error
		e, err = casbin.NewEnforcer(path.Join(modelPath, "model.conf"), path.Join(policyPath, "/policy.csv"))
		if err != nil {
			logging.Fatal(err)
		}
		break
	default:
		logging.Fatal("Wrong config to casbin policy")
		break
	}
	err := e.LoadPolicy()
	if err != nil {
		logging.Error(err)
	}
	enforcer = e
}

func GetEnforcer() *casbin.Enforcer {
	return enforcer
}
