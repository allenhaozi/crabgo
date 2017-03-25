package crabgo

var RouterList map[string]ControllerInterface

func init() {
	RouterList = make(map[string]ControllerInterface)
}

func Router(key string, ctrl ControllerInterface) {
	if _, ok := RouterList[key]; !ok {
		RouterList[key] = ctrl
	}
}
