package mesh

import (
	"github.com/allenhaozi/crabgo/pkg/register"
)

type LicenseService struct{}

func NewLicenseService(cfg *register.Config) *LicenseService {
	return &LicenseService{}
}
