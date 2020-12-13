package server

import (
	"net/http"

	"github.com/allenhaozi/crabgo/pkg/apis/common"

	sageruntime "github.com/allenhaozi/crabgo/pkg/register/runtime"
	echo "github.com/labstack/echo/v4"
	alog "github.com/sirupsen/logrus"

	"github.com/allenhaozi/crabgo/pkg/register"
)

type WebServer struct {
	*echo.Echo
}

func NewWebServer() *WebServer {
	ins := &WebServer{}
	ins.Echo = echo.New()
	return ins
}

// start a http web server
func (ws *WebServer) StartHttpServer(cfg *register.Config) error {

	alog.Infof("success start web server port:%s", cfg.GeneralConfig.CrabConfig.WebServerPort)
	//register http health check
	ws.patchHealthCheck()
	// start server at the end of the code
	err := ws.Start(":" + cfg.GeneralConfig.CrabConfig.WebServerPort)
	return err
}

//register restful api support path list
// methdo + path -> handler function
func (ws *WebServer) RegisterHttpHandler(apiSet map[sageruntime.GroupVersion][]common.RestfulApiMeta) {
	for gv, apiMetaList := range apiSet {
		for _, apiMeta := range apiMetaList {
			prefix := gv.GetGroupVersionPath()
			path := prefix + apiMeta.Path
			ws.Add(
				apiMeta.Method,
				path,
				apiMeta.HandlerFunc,
			)
			alog.Infof("register handler func, method:%s,path:%s", apiMeta.Method, path)
		}
	}
}

func (ws *WebServer) patchHealthCheck() {
	healthCheck := func(c echo.Context) error {
		return c.String(200, "ok")
	}
	ws.Add(http.MethodGet, "/health", healthCheck)
}
