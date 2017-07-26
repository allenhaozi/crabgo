package crabgo

import (
	"net/http"
	"reflect"
)

type CrabHandler struct {
	*CrabRouter
}

func NewCrabHandler() *CrabHandler {
	handlers := &CrabHandler{
		&CrabRouter{RouterList: make(map[string]ControllerInterface)},
	}
	return handlers
}

func (self *CrabHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// parse the controller name and action name
	//execute requst by the relevant controller and action
	dispatcher := NewCrabDispatcher()
	dispatcher.Dispatch(rw, req)

	// parse request parameter
	parameter := new(CrabParameter)
	parameter.ParseParameter(req)
	//execute requst by the relevant controller and action
	if controller, ok := self.RouterList[dispatcher.ControllerNameKey]; ok {
		//initialize controller data
		controller.Init(rw, req)
		//inject request parameter to controller struct
		for k, v := range parameter.RetParams {
			controller.SetParam(k, v)
		}
		ctrlType := reflect.TypeOf(controller)
		//todo  reflect: Call with too few input arguments
		if _, ok := ctrlType.MethodByName(dispatcher.ActionName); ok {
			//in := make([]reflect.Value, 0)
			//method.Func.Call(in)
			value := reflect.ValueOf(controller)
			//execute controller action(a function)
			value.MethodByName(dispatcher.ActionName).Call(nil)
		} else {
			panic(ErrNotFoundAction)
		}
	} else {
		panic(ErrNotFoundCtrl)
	}
}
