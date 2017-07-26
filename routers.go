package crabgo

type CrabRouter struct {
	RouterList map[string]ControllerInterface
}

func (self *CrabRouter) AddRouter(ctrlNameKey string, ctrl ControllerInterface) {
	if _, ok := self.RouterList[ctrlNameKey]; !ok {
		self.RouterList[ctrlNameKey] = ctrl
	}
}
