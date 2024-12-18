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

package main

import (
	"context"
)

import (
	"dubbo.apache.org/dubbo-go/v3/client"
	_ "dubbo.apache.org/dubbo-go/v3/imports"

	"github.com/dubbogo/gost/log/logger"
)

import (
	greet "github.com/apache/dubbo-go-samples/java_interop/protobuf-triple/go/proto"
)

func main() {

	cli, err := client.NewClient(
		client.WithClientURL("127.0.0.1:50052"),
	)
	if err != nil {
		panic(err) // If there's an error, it's handled immediately by panicking.
	}

	svc, err := greet.NewGreeter(cli)
	if err != nil {
		panic(err) // Same here, handle the error immediately.
	}

	resp, err := svc.SayHello(context.Background(), &greet.HelloRequest{Name: "hello world"})
	if err != nil {
		logger.Error(err)
		return // Now, we explicitly handle the error by logging it and then returning from the function.
	}
	logger.Infof("Greet response: %s", resp.Message)
}
