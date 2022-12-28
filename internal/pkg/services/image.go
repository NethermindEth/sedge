package services

import "context"

func (s *serviceManager) Image(ct string) (string, error) {
	info, err := s.dockerClient.ContainerInspect(context.Background(), ct)
	if err != nil {
		return "", err
	}
	return info.Image, nil
}
