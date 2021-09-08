// Copyright (C) 2015 Nippon Telegraph and Telephone Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package glib

import (
	"fmt"
	"os"

	"git.netdike/backend/gobgp2/internal/pkg/version"
	"google.golang.org/grpc"
)

func glib() {
	if len(os.Args) > 1 && os.Args[1] == "--version" {
		fmt.Println("gobgp version", version.Version())
		os.Exit(0)
	}
	grpc.EnableTracing = false
	newRootCmd().Execute()
}