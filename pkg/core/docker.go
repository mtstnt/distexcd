package core

import (
	"fmt"

	"github.com/docker/docker/client"
)

var dcl *client.Client

func ConfigureDockerClient() error {
	c, err := client.NewClientWithOpts(
		client.WithHostFromEnv(),
		client.WithAPIVersionNegotiation(),
	)
	if err != nil {
		return fmt.Errorf("error getting docker client: %w", err)
	}
	dcl = c
	return nil
}

func GetDockerClient() *client.Client {
	if dcl == nil {
		GetLogger().Fatal("Docker client is not initialized!")
	}
	return dcl
}
