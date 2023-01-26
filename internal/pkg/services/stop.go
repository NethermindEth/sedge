/*
Copyright 2022 Nethermind

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
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
		log.Infof("Stopping service: %s, currently on %s status", service, ctInfo.State.Status)
		timeout := 5 * time.Minute
		if err := s.dockerClient.ContainerStop(context.Background(), ctInfo.ID, &timeout); err != nil {
			return fmt.Errorf("%w %s: %s", ErrStoppingContainer, service, err)
		}
	}
	return nil
}
