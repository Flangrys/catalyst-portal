package docker

import (
	"context"
	"log"

	"github.com/docker/docker/client"
)

var DockerCli *client.Client
var DockerCtx context.Context

func Init() {
	var err error

	DockerCtx = context.Background()
	DockerCli, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	if err != nil {
		log.Panic("Docker just having a bad day.")
	}

	defer DockerCli.Close()
}
