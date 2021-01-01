/*
 *    Copyright 2019 Samsung SDS
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const configFIle = "../../testdata/accelerator_test.yaml"

func TestConfig(t *testing.T) {
	conf, err := loadConfig(configFIle)
	assert.NoError(t, err)
	assert.Equal(t, "localhost", conf.Host)
	assert.Equal(t, 5050, conf.Port)
	assert.Equal(t, "Admin", conf.UserName)
	assert.Equal(t, "ping", conf.Batch[0]["fcn"])
	assert.Equal(t, 10, conf.Batch[1]["maxBatchItems"])

	client, err := conf.BatchClient()
	assert.NoError(t, err)
	assert.NotNil(t, client)
}
