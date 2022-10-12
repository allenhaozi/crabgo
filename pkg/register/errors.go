/**
 * Copyright 2021 4Paradigm
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package register

import (
	"fmt"
	"net/http"
	"strconv"
)

type CrabError struct {
	HttpCode     int              `json:"-"`
	ResponseType CrabResponseType `json:"-"`
	Message      string           `json:"msg"`
	Reason       string           `json:"reason,omitempty"`
	Status       string           `json:"status"`
}

var errMsgMap = map[int]string{
	http.StatusOK:            "success",
	http.StatusNotAcceptable: "request not acceptable",
	http.StatusNotFound:      "request resource not found",
	http.StatusForbidden:     "request forbidden",

	http.StatusBadRequest: "request failure",

	http.StatusInternalServerError: "response controller failure",
}

func NewCrabError() *CrabError {
	return &CrabError{
		HttpCode:     http.StatusOK,
		ResponseType: CrabResponseTypeJSON,
		Message:      errMsgMap[http.StatusOK],
		Status:       "0",
	}

}

// implement interface CrabResponseInf
func (i *CrabError) SetData(code int, t CrabResponseType, dataList ...interface{}) {

	i.ResponseType = t

	if msg, ok := errMsgMap[code]; ok {
		i.Message = msg
	} else {
		i.Message = errMsgMap[http.StatusBadRequest]
	}
	i.HttpCode = code
	i.Status = strconv.Itoa(code)

	var msg string

	if len(dataList) > 0 {
		for _, item := range dataList {
			if str, ok := item.(string); ok {
				str := str
				if len(msg) == 0 {
					msg = str
				} else {
					msg = fmt.Sprintf("%s;%s", msg, str)
				}
			}
		}
		i.Reason = msg
	}
}

func (i *CrabError) GetMsg(code int) string {
	if str, ok := errMsgMap[code]; ok {
		return str
	}
	return errMsgMap[http.StatusBadRequest]
}

func (i *CrabError) Success() bool {
	return i.HttpCode == http.StatusOK
}
