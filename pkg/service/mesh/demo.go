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

package mesh

import (
	"github.com/pkg/errors"

	crabcorev1 "github.com/allenhaozi/crabgo/apis/core/v1"
	"github.com/allenhaozi/crabgo/pkg/register"
)

type DemoService struct {
	cfg *register.Config
}

func NewDemoService(cfg *register.Config) *DemoService {
	return &DemoService{cfg: cfg}
}

func (s *DemoService) GetUser(ctx register.Context, userId string) (*crabcorev1.User, error) {
	if len(userId) == 0 {
		return nil, errors.New("not found user id")
	}

	user := &crabcorev1.User{}
	user.Name = "demo"
	user.UserId = userId

	return user, nil
}
