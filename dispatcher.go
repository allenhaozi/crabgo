package crabgo

import (
	"github.com/Allenhaozi/crabgo/utils"
	"net/http"
	"strings"
)

const (
	// if your url is :  http://127.0.0.1:8080/user/info
	// this framework will parse the uri [ /user/info ]
	// the first string is [ user ] : will be controller prefix, the controller name is : [ UserController ]
	// the second string is [ info ] : will be action prefix, the action name is : [ InfoAction]
	//
	// caution : only first two string will be used
	uriLayerNum int = 2
	//default action IndexAction
	defAction string = "Index"
	//default action IndexAction
	actionSuffix string = "Action"
	//controller name sufix
	ctrlSuffix string = "Controller"
)

var (
	CrabDispatcher *Dispatcher
)

type Dispatcher struct {
	ControllerNameKey string
	ActionName        string
}

func NewCrabDispatcher() *Dispatcher {
	return new(Dispatcher)
}

func (self *Dispatcher) Dispatch(rw http.ResponseWriter, req *http.Request) {
	ctrlNameKey, actionName := getCtrlAndAction(req)
	//this is the controller map key , this map is in routers.go
	self.ControllerNameKey = ctrlNameKey
	// first character upper
	self.ActionName = actionName

}

func getCtrlAndAction(req *http.Request) (ctrlNameKey string, actionName string) {
	uri := req.URL.Path
	//todo sepcial characters process
	slicePath := strings.Split(uri, "/")
	if len(slicePath) > uriLayerNum {
		ctrlNameKey = slicePath[1]
		actionName = slicePath[2]
	} else if len(slicePath) == 2 {
		ctrlNameKey = slicePath[1]
		actionName = defAction
	} else {
		ctrlNameKey = defAction
		actionName = defAction
	}

	//ensure string is lowercase
	//this is the controller map key , this map is in routers.go
	ctrlNameKey = strings.ToLower(ctrlNameKey)
	// first character upper
	actionName = strings.ToLower(actionName)
	actionName = utils.UcFirst(strings.ToLower(actionName)) + actionSuffix

	return
}
