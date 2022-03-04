package main

import (
	"fmt"

	"github.com/heroku/docker-registry-client/registry"
	"github.com/stackrox/rox/pkg/utils"
)

func main() {
	doDockerHub()
}

func doDockerHub() {
	reg, err := registry.New("https://registry-1.docker.io", "roxross", "Docker@StackRox")
	utils.CrashOnError(err)

	digest, _, err := reg.ManifestDigest("centos", "8")
	utils.CrashOnError(err)

	fmt.Println(digest)
}
//
//func doRedHat() {
//	redhatRegistry, err := registry.New("https://registry.access.redhat.com")
//	utils.CrashOnError(err)
//
//
//}
