package options

import (
	"github.com/allenhaozi/crabgo/pkg/register"
)

type Options struct {
	Config *register.Config
}

func NewOptions() (*Options, error) {

	o := &Options{}
	var err error
	o.Config, err = register.NewConfig()
	if err != nil {
		return nil, err
	}

	return o, err
}
