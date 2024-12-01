package docker

import (
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
)

// Get a list of all running containers on this host and a count of these.
func GetRunningContainers() ([]types.Container, error) {
	containers, err := DockerCli.ContainerList(DockerCtx, container.ListOptions{
		Filters: filters.NewArgs(filters.Arg("status", "running")),
	})

	if err != nil {
		return nil, fmt.Errorf("[containers] cannot get running containers: %s", err)
	}

	return containers, nil
}

// Get a list of containers which matches with the given ID.
func GetDetailedContainer(id string) ([]types.Container, error) {

	if id == "" {
		return nil, fmt.Errorf("[containers] cannot provide an empty container id")
	}

	containers, err := DockerCli.ContainerList(DockerCtx, container.ListOptions{
		Limit:   1,
		Filters: filters.NewArgs(filters.Arg("id", id)),
	})

	if err != nil {
		return nil, fmt.Errorf("[containers] cannot retrieve the given container with the id: %s", id[:10])
	}

	return containers, nil
}
