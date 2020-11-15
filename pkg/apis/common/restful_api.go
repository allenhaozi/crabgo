package common

import "github.com/labstack/echo/v4"

type RestfulApiMeta struct {
	Mod         string
	Method      string
	Act         string
	Path        string
	HandlerFunc func(e echo.Context) error
}

func (ra RestfulApiMeta) Gen(mod, m, a, p string) RestfulApiMeta {
	ra.Mod = mod
	ra.Method = m
	ra.Act = a
	ra.Path = p
	return ra
}
