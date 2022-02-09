// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package client

import (
	"context"
	"sync"
	"testing"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"

	"github.com/cloudwego/kitex-examples/hello/kitex_gen/api"
	"github.com/cloudwego/kitex-examples/hello/kitex_gen/api/hello"
)

func TestClient(t *testing.T) {
	client, err := hello.NewClient(
		"hello", client.WithHostPorts("0.0.0.0:19999"),
		client.WithBackupRequest(retry.NewBackupPolicy(10)),
	)
	if err != nil {
		t.Fatalf("%v", err)
	}
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := client.Echo(context.Background(), &api.Request{Message: "my request"})
			if err != nil {
				t.Fatalf("%v", err)
			}
		}()
	}
	wg.Wait()
}
