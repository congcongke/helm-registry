/*
Copyright 2017 congcongke authors. All rights reserved.
*/

package main

import (
	"github.com/congcongke/helm-registry/cmd/registry/cmd"
	_ "github.com/congcongke/helm-registry/pkg/storage/simple"
	_ "github.com/docker/distribution/registry/storage/driver/filesystem"
)

func main() {
	cmd.Run()
}
