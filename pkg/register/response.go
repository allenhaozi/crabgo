/**
 * Copyright 2021 Crabgo Authors
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
	"net/http"

	"github.com/pkg/errors"

	"github.com/allenhaozi/crabgo/pkg/utils"
)

type CrabResponseType string

var (
	CrabResponseTypeRedirect CrabResponseType = "Redirect"
	CrabResponseTypeJSON     CrabResponseType = "JSON"
)

type CrabResponse struct {
	HttpCode     int              `json:"-"`
	Msg          string           `json:"msg"`
	Status       string           `json:"status"`
	ResponseType CrabResponseType `json:"-"`
	Data         interface{}      `json:"data" swaggertype:"array,object"`
}

func NewCrabResponse() *CrabResponse {
	return &CrabResponse{
		HttpCode: http.StatusOK,
	}
}

// implement interface CrabResponseInf
func (ir *CrabResponse) SetData(code int, t CrabResponseType, dataList ...interface{}) {
	ir.HttpCode = code
	ir.ResponseType = t

	if len(dataList) >= 1 {
		ir.Msg = "success"
		ir.Status = "0"
		ir.Data = dataList[0]
	}
}

func (ir *CrabResponse) GetRedirectURL() (string, error) {
	if url, ok := ir.Data.(string); ok {
		if utils.IsURL(url) {
			return url, nil
		}
	}
	return "", errors.New("redirect url can not be found")
}

func IsValidResponse(t string) error {
	switch t {
	case string(CrabResponseTypeJSON):
		return nil
	case string(CrabResponseTypeRedirect):
		return nil
	default:
		return errors.Errorf("invalid response type, %s is not supported", t)
	}
}
