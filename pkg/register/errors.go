package register

import (
	"net/http"
	"strconv"
)

type SageError struct {
	HttpCode int    `json:"-"`
	Message  string `json:"msg"`
	Reason   string `json:"reason,omitempty"`
	Status   string `json:"status"`
}

var errMsgMap = map[int]string{
	http.StatusOK:            "success",
	http.StatusNotAcceptable: "request not acceptable",
	http.StatusNotFound:      "request resource not found",
	http.StatusForbidden:     "request forbidden",

	http.StatusBadRequest: "request failure",

	http.StatusInternalServerError: "response controller failure",
}

func NewSageError() *SageError {
	return &SageError{
		HttpCode: http.StatusOK,
		Message:  errMsgMap[http.StatusOK],
		Status:   "0",
	}

}

// implement interface ProphetResponseInf
func (se *SageError) SetData(code int, dataList ...interface{}) {
	if msg, ok := errMsgMap[code]; ok {
		se.Message = msg
	} else {
		se.Message = errMsgMap[http.StatusBadRequest]
	}
	se.HttpCode = code
	se.Status = strconv.Itoa(code)
	//TODO one reason args can satisfy our demand
	if len(dataList) > 0 {
		if str, ok := dataList[0].(string); ok {
			se.Reason = str
		}
	}
}

func (se *SageError) GetMsg(code int) string {
	if str, ok := errMsgMap[code]; ok {
		return str
	}
	return errMsgMap[http.StatusBadRequest]
}

func (se *SageError) Succ() bool {
	if se.HttpCode == http.StatusOK {
		return true
	}
	return false
}
