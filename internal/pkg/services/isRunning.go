package services

import "context"

func (s *serviceManager) IsRunning(ct string) (bool, error) {
	info, err := s.dockerClient.ContainerInspect(context.Background(), ct)
	if err != nil {
		return false, err
	}
	return info.State.Running, nil
}
