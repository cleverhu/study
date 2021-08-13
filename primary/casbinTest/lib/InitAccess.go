package lib

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"

	"log"
)

var E *casbin.Enforcer

func init() {
	DB = initGormDB()
	E = initE()
	initPolicy()
}

func initE() *casbin.Enforcer {
	adapter, err := gormadapter.NewAdapterByDB(DB)
	if err != nil {
		log.Fatal(err)
	}
	e, err := casbin.NewEnforcer("resources/model.conf", adapter)
	if err != nil {
		log.Fatal(err)
	}
	return e
}

func initPolicy() {
	initRoles()
	initUserRoles()
	initRouteRolesPolicy()
}
