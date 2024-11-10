package docker

import (
	"log"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
)

type DummyContainer struct {
	Id     string `json:"Id"`
	Name   string `json:"Name"`
	Age    int64  `json:"Age"`
	Status string `json:"Status"`
}

func GetRunningContainers() ([]DummyContainer, int) {
	containers, err := DockerCli.ContainerList(DockerCtx, container.ListOptions{
		Filters: filters.NewArgs(filters.Arg("status", "running")),
	})

	if err != nil {
		log.Fatalf("[Containers] Cannot get running containers: %s", err)
		return nil, 0
	}

	var dummys []DummyContainer

	for _, container := range containers {
		dummys = append(dummys, DummyContainer{
			Name:   container.Names[0],
			Id:     container.ID,
			Age:    container.Created,
			Status: container.Status,
		})
	}

	return dummys, len(containers)
}
