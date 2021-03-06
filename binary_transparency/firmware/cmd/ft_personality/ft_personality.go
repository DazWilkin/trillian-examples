// Copyright 2020 Google LLC. All Rights Reserved.
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

// This package is the entrypoint for the Firmware Transparency personality server.
// This requires a Trillian instance to be reachable via gRPC and a tree to have
// been provisioned. See the README in the root of this project for instructions.
// Start the server using:
// go run ./cmd/ft_personality/main.go --logtostderr -v=2 --tree_id=$TREE_ID
package main

import (
	"context"
	"flag"
	"time"

	"github.com/golang/glog"
	"github.com/google/trillian-examples/binary_transparency/firmware/cmd/ft_personality/impl"
)

var (
	listenAddr = flag.String("listen", ":8000", "address:port to listen for requests on")

	connectTimeout = flag.Duration("connect_timeout", time.Second, "the timeout for connecting to the backend")
	trillianAddr   = flag.String("trillian", ":8090", "address:port of Trillian Log gRPC service")
	treeID         = flag.Int64("tree_id", -1, "the tree ID of the log to use")

	casDBFile = flag.String("cas_db_file", "", "Path to a file to be used as sqlite3 storage for images, e.g. /tmp/ft.db")

	sthRefresh = flag.Duration("sth_refresh_interval", 5*time.Second, "how often to fetch the latest log root from Trillian")
)

func main() {
	flag.Parse()

	ctx := context.Background()
	if err := impl.Main(ctx, impl.PersonalityOpts{
		ListenAddr:     *listenAddr,
		ConnectTimeout: *connectTimeout,
		TrillianAddr:   *trillianAddr,
		TreeID:         *treeID,
		CASFile:        *casDBFile,
		STHRefresh:     *sthRefresh,
	}); err != nil {
		glog.Exit(err.Error())
	}
}
