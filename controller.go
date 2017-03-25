package crabgo

type ControllerInterface interface {
	IndexAction()
	Init()
	SetParam(string, string)
	GetParam(string) string
}

type Controller struct {
	RetParam map[string]string
	Response map[string]interface{}
}

func (self *Controller) Init() {
	self.RetParam = make(map[string]string)
	self.Response = make(map[string]interface{})
}

func (self *Controller) SetResponse(key string, data interface{}) {
	self.Response[key] = data
}

func (self *Controller) GetResponse(key string) interface{} {
	if data, ok := self.RetParam[key]; ok {
		return data
	} else {
		return nil
	}
}

func (self *Controller) SetParam(key string, value string) {
	self.RetParam[key] = value
}

func (self *Controller) GetParam(key string) string {
	if data, ok := self.RetParam[key]; ok {
		return data
	} else {
		return ""
	}
}

func (self *Controller) GetAllParam() map[string]string {
	return self.RetParam
}

func (self *Controller) SetAllParam(param map[string]string) {
	self.RetParam = param
}
