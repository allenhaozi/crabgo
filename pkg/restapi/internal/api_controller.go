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

package internal

import (
	"github.com/allenhaozi/crabgo/pkg/register"
)

type ApiController struct {
	cfg *register.Config
}

func (ac *ApiController) PreExecute(ctx *register.Context) register.CrabResponseIf {
	return ac.preCheck(ctx)
}

func (ac *ApiController) preCheck(ctx *register.Context) *register.CrabError {
	errInfo := register.NewCrabError()
	return errInfo
}

func (ac *ApiController) PostExecute(ctx *register.Context) register.CrabResponseIf {
	return register.NewCrabError()
}

func (ac *ApiController) Initial(cfg *register.Config) {
	ac.cfg = cfg
}
