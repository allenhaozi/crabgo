package register

import (
	"net/http"
)

type SageResponseIf interface {
	SetData(code int, dataList ...interface{})
}

type SageResponse struct {
	HttpCode int         `json:"-"`
	Msg      string      `json:"msg"`
	Status   string      `json:"status"`
	Data     interface{} `json:"data"`
}

func NewSageResponse() *SageResponse {
	return &SageResponse{
		HttpCode: http.StatusOK,
	}
}

//implement interface ProphetResponseInf
func (sr *SageResponse) SetData(code int, dataList ...interface{}) {
	sr.HttpCode = code
	if len(dataList) >= 1 {
		sr.Msg = "success"
		sr.Status = "0"
		sr.Data = dataList[0]
	}
}
