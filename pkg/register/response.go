package register

import (
	"net/http"
)

type CrabResponseIf interface {
	SetData(code int, dataList ...interface{})
}

type CrabResponse struct {
	HttpCode int         `json:"-"`
	Msg      string      `json:"msg"`
	Status   string      `json:"status"`
	Data     interface{} `json:"data"`
}

func NewCrabResponse() *CrabResponse {
	return &CrabResponse{
		HttpCode: http.StatusOK,
	}
}

//implement interface ProphetResponseInf
func (sr *CrabResponse) SetData(code int, dataList ...interface{}) {
	sr.HttpCode = code
	if len(dataList) >= 1 {
		sr.Msg = "success"
		sr.Status = "0"
		sr.Data = dataList[0]
	}
}
