// copied from: https://github.com/istio/istio/blob/master/mixer/tools/mixgen/main.go

// Copyright 2016 Istio Authors
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

package main

import (
	"os"

	"istio.io/istio/mixer/cmd/shared"
	"istio.io/istio/mixer/tools/mixgen/cmd"
	"istio.io/pkg/version"
)

func main() {
	rootCmd := cmd.GetRootCmd(os.Args[1:], shared.Printf, shared.Fatalf)

	rootCmd.AddCommand(version.CobraCommand())

	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}