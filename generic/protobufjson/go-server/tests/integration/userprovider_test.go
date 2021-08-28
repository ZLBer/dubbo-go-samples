// +build integration

/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package integration

import (
	"context"
	"testing"
)

import (
	"dubbo.apache.org/dubbo-go/v3/config"

	hessian "github.com/apache/dubbo-go-hessian2"

	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	o, err := referenceConfig.GetRPCService().(*config.GenericService).Invoke(
		context.TODO(),
		[]interface{}{
			"GetUser",
			[]string{},
			[]hessian.Object{"{\"id\":1}"},
		},
	)
	assert.Nil(t, err)
	resp := o.(string)
	assert.Equal(t, `{"name":"XavierNiu","age":"20"}`, resp)
}