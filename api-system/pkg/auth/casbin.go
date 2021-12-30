package auth

import (
	"axiangcoding/antonstar/api-system/internal/app/data"
	"axiangcoding/antonstar/api-system/pkg/logging"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

var enforcer *casbin.Enforcer

func SetupCasbin() {
	// Use gormadapter to load the policy from the database, and use an existing gorm instance
	// 使用gormadapter来从数据库中加载policy，并且使用的是已经存在的gorm实例
	adapter, err := gormadapter.NewAdapterByDB(data.GetDB())
	if err != nil {
		logging.Error(err)
	}
	e, err := casbin.NewEnforcer("config/default/model.conf", adapter)
	if err != nil {
		logging.Error(err)
	}
	err = e.LoadPolicy()
	if err != nil {
		logging.Error(err)
	}
	enforcer = e
}

func GetEnforcer() *casbin.Enforcer {
	return enforcer
}
