package services

import (
	"context"
	"io/ioutil"

	"github.com/docker/docker/api/types"
	log "github.com/sirupsen/logrus"
)

// TODO: Write unit tests for this function
// ContainerLogs returns the logs of a container. <service> is the name of the container for logging purposes.
func (s *serviceManager) ContainerLogs(ctID, service string) (string, error) {
	logReader, err := s.dockerClient.ContainerLogs(context.Background(), ctID, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Follow:     false,
	})
	if err != nil {
		return "", err
	}
	defer logReader.Close()
	logs, err := ioutil.ReadAll(logReader)
	if err == nil {
		log.Debugf("%s container logs: %s", service, string(logs))
	}
	return string(logs), err
}
