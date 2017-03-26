package crabgo

import (
	"github.com/Allenhaozi/crabgo/utils"
	"net/http"
	"reflect"
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
	ControllerNameKey string
	ActionName        string
}

func (self *CrabDispatcher) Dispatch(rw http.ResponseWriter, req *http.Request) {
	ctrlNameKey, actionName := getCtrlAndAction(req)
	//this is the controller map key , this map is in routers.go
	self.ControllerNameKey = ctrlNameKey
	// first character upper
	self.ActionName = actionName

	// parse request parameter
	parameter := new(CrabParameter)
	parameter.ParseParameter(req)
	//execute requst by the relevant controller and action
	if controller, ok := RouterList[ctrlNameKey]; ok {
		//initialize controller data
		controller.Init(rw, req)
		//inject request parameter to controller struct
		for k, v := range parameter.RetParams {
			controller.SetParam(k, v)
		}
		value := reflect.ValueOf(controller)
		//execute controller action(a function)
		value.MethodByName(actionName).Call(nil)
	} else {
		panic(ErrNotFound)
	}
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
