package services

import (
	"context"
	"fmt"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	log "github.com/sirupsen/logrus"
)

const (
	ServiceValidator = "validator"
)

var ServiceContainer map[string]string = map[string]string{
	ServiceValidator: "validator-client",
}

func Image(dockerClient client.APIClient, serviceName string) (string, error) {
	info, err := ctInfo(dockerClient, serviceName)
	return info.Image, err
}

func Stop(serviceName string) error {
	containerName, ok := ServiceContainer[serviceName]
	if !ok {
		return fmt.Errorf("unknown container name for service: %s", serviceName)
	}
	dockerCli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return err
	}
	defer dockerCli.Close()
	ctInfo, err := dockerCli.ContainerInspect(context.Background(), containerName)
	if err != nil {
		if client.IsErrNotFound(err) {
			return nil
		}
		return err
	}
	if ctInfo.State.Running {
		log.Infof("stopping service: %s, currently on %s status", serviceName, ctInfo.State.Status)
		timeout := 5 * time.Minute
		if err := dockerCli.ContainerStop(context.Background(), containerName, &timeout); err != nil {
			return fmt.Errorf("error stopping service %s (with container %s): %w", serviceName, containerName, err)
		}
	}
	return nil
}

func Start(serviceName string) error {
	containerName, ok := ServiceContainer[serviceName]
	if !ok {
		return fmt.Errorf("unknown container name for service: %s", serviceName)
	}
	dockerCli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return err
	}
	defer dockerCli.Close()
	if err := dockerCli.ContainerStart(context.Background(), containerName, types.ContainerStartOptions{}); err != nil {
		return fmt.Errorf("error starting service %s (with container %s): %w", serviceName, containerName, err)
	}
	return nil
}

func Volumes(dockerClient client.APIClient, serviceName string) (map[string]struct{}, error) {
	info, err := ctInfo(dockerClient, serviceName)
	return info.Config.Volumes, err
}

func IsRunning(serviceName string) (bool, error) {
	dockerCli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return false, err
	}
	defer dockerCli.Close()
	info, err := ctInfo(dockerCli, serviceName)
	return info.State.Running, err
}
