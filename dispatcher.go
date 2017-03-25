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

type CrabDispatcher struct {
	ContrllerNameKey string
	ActionName       string
}

func (self *CrabDispatcher) Dispatch(req *http.Request) {
	var ctrl string
	var action string
	uri := req.URL.Path
	//todo sepcial characters process
	slicePath := strings.Split(uri, "/")
	if len(slicePath) > uriLayerNum {
		ctrl = slicePath[1]
		action = slicePath[2]
	} else if len(slicePath) == 2 {
		ctrl = slicePath[1]
		action = defAction
	} else {
		ctrl = defAction
		action = defAction
	}

	//ensure string is lowercase
	ctrl = strings.ToLower(ctrl)
	action = strings.ToLower(action)

	//this is the controller map key , this map is in routers.go
	self.ContrllerNameKey = ctrl
	// first character upper
	self.ActionName = utils.UcFirst(action) + actionSuffix
}
