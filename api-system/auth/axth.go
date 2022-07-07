package auth

import (
	"github.com/axiangcoding/ax-web/data"
	"github.com/axiangcoding/ax-web/logging"
	"github.com/axiangcoding/axth"
)

var axthEnforcer *axth.Enforcer

func SetupAxth() {
	options, err := axth.DefaultOptions()
	if err != nil {
		logging.Fatal(err)
	}
	e, err := axth.NewEnforcer(data.GetDB(), options)
	if err != nil {
		logging.Fatal(err)
	}
	axthEnforcer = e
}

func GetAxthEnforcer() *axth.Enforcer {
	return axthEnforcer
}
