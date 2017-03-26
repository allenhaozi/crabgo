package crabgo

import (
	"errors"
)

var ErrInvalidParam error = errors.New("invalid parameter")
var ErrNotFound error = errors.New("controller not found")
