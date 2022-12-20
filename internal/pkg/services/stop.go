package services

import (
	"context"
	"fmt"
	"time"

	"github.com/docker/docker/client"
	log "github.com/sirupsen/logrus"
)

func (s *serviceManager) Stop(service string) error {
	ctInfo, err := s.dockerClient.ContainerInspect(context.Background(), service)
	if err != nil {
		if client.IsErrNotFound(err) {
			return nil
		}
		return err
	}
	if ctInfo.State.Running {
		log.Infof("stopping service: %s, currently on %s status", service, ctInfo.State.Status)
		timeout := 5 * time.Minute
		if err := s.dockerClient.ContainerStop(context.Background(), service, &timeout); err != nil {
			return fmt.Errorf("%w %s: %s", ErrStoppingContainer, service, err)
		}
	}
	return nil
}
