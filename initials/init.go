package initials

import (
	"github.com/hardbornman/garglewool-service/dao"
	"github.com/hardbornman/garglewool-service/initials/config"
)

func init() {
	initLog()
	config.InitConfig()
	initApp()
	initId()
	dao.Init()
}
