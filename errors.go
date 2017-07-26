package crabgo

import (
	"errors"
)

var ErrInvalidParam error = errors.New("invalid parameter")
var ErrNotFoundCtrl error = errors.New("controller not found")
var ErrNotFoundAction error = errors.New("aciton not found")
