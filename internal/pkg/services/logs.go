package services

import (
	"bytes"
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/pkg/stdcopy"
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

	var logs bytes.Buffer
	written, err := stdcopy.StdCopy(&logs, &logs, logReader)
	log.Debug("Logs written %d bytes", written)
	if err == nil {
		log.Debugf("%s container logs: %s", service, logs.String())
	}
	return logs.String(), err
}
