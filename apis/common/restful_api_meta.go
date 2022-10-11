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

package common

import (
	"fmt"
	"regexp"

	"github.com/labstack/echo/v4"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var pRegex = regexp.MustCompile("^/")

type RestAPIMeta struct {
	ApiGroup     ApiGroup
	Handler      string
	Method       string
	Act          string
	Path         string
	GroupVersion schema.GroupVersion
	HandlerFunc  func(e echo.Context) error
}

func (r RestAPIMeta) Gen(gv schema.GroupVersion, apiGroup ApiGroup, mod, m, a, p string) RestAPIMeta {

	r.ApiGroup = apiGroup
	r.Handler = mod
	r.Method = m
	r.Act = a
	r.Path = p
	r.GroupVersion = gv

	return r
}

// compose api path
func (r RestAPIMeta) ApiPath() string {

	if !pRegex.MatchString(r.Path) {
		r.Path = "/" + r.Path
	}
	return fmt.Sprintf("%s/%s%s", r.ApiGroup, r.GroupVersion.Version, r.Path)
}
