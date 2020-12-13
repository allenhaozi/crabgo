package restfulapi

import (
	"github.com/allenhaozi/crabgo/pkg/register"
)

type ApiController struct {
	cfg *register.Config
}

func (ac *ApiController) PreExecute(ctx *register.Context) register.CrabResponseIf {
	// assign config to context
	ctx.Config = ac.cfg

	return ac.preCheck(ctx)
}

func (ac *ApiController) preCheck(ctx *register.Context) *register.CrabError {
	return ac.checkPermission(ctx)
}

// check user request permission
func (ac *ApiController) checkPermission(ctx *register.Context) *register.CrabError {
	errInfo := register.NewCrabError()
	return errInfo
}

func (ac *ApiController) PostExecute(ctx *register.Context) register.CrabResponseIf {
	errInfo := register.NewCrabError()
	return errInfo
}

func (ac *ApiController) initial(cfg *register.Config) {
	ac.cfg = cfg
}
