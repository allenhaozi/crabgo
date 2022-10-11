package server

import (
	"io"
	"net/http"

	echo "github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	echoswagger "github.com/swaggo/echo-swagger"
	"gitlab.4pd.io/openaios/openaios-iam/apis/common"
	"gitlab.4pd.io/openaios/openaios-iam/pkg/register"
	"k8s.io/apimachinery/pkg/runtime/schema"
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

	log.Infof("success start web server port:%s", cfg.GeneralConfig.IAMConfig.WebServerPort)
	// register http health check
	ws.patchHealthCheck()
	// swagger handler for api docs
	ws.patchSwaggerHandler()
	// start server at the end of the code
	err := ws.Start(":" + cfg.GeneralConfig.IAMConfig.WebServerPort)
	return err
}

//register restful api support path list
// method + path -> handler function
func (ws *WebServer) RegisterHttpHandler(apiSet map[schema.GroupVersion][]common.RestAPIMeta) {

	for _, apiMetaList := range apiSet {
		for _, apiMeta := range apiMetaList {
			path := apiMeta.ApiPath()
			ws.Add(
				apiMeta.Method,
				path,
				apiMeta.HandlerFunc,
			)
			log.Infof("register handler func, method: %s, path: %s", apiMeta.Method, path)
		}
	}

	// register not found handler
	ws.registerNotFoundHandler()
}

func (ws *WebServer) registerNotFoundHandler() {
	echo.NotFoundHandler = func(ctx echo.Context) error {
		r := ctx.Request()
		log.Warnf("request not found \n path: %s \n method: %s \n query: %s \n", r.URL.Path[1:], r.Method, r.URL.RawQuery)
		if b, err := io.ReadAll(r.Body); err == nil {
			log.Warnf("request not found, body: \n %s ", string(b))
		}
		errInfo := register.NewCrabError()
		errInfo.SetData(http.StatusNotFound, "request failed")
		return ctx.JSON(http.StatusNotFound, errInfo)
	}
}

func (ws *WebServer) patchHealthCheck() {
	healthCheck := func(c echo.Context) error {
		return c.String(200, "ok")
	}
	ws.Add(http.MethodGet, "/health", healthCheck)
}

func (ws *WebServer) patchSwaggerHandler() {
	ws.Add(http.MethodGet, "/swagger/*", echoswagger.WrapHandler)
}
