package health

import (
	"github.com/dangduoc08/ginject/common"
	"github.com/dangduoc08/ginject/core"
)

type HealthController struct {
	common.REST
}

func (instance HealthController) NewController() core.Controller {

	return instance
}

func (instance HealthController) READ_VERSION_1() string {

	return "Hello API"
}
