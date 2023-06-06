package clientsimages

import (
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

const (
	OnlineImagesSourceFile = "https://raw.githubusercontent.com/NethermindEth/sedge/develop/configs/client_images.yaml"
)

type OnlineClientsImages struct {
	ClientsImages DefaultClientsImages
}

func (oci OnlineClientsImages) Execution() ExecutionClientsImages {
	return oci.ClientsImages.Execution()
}

func (oci OnlineClientsImages) Consensus() ConsensusClientsImages {
	return oci.ClientsImages.Consensus()
}

func (oci OnlineClientsImages) Validator() ValidatorClientsImages {
	return oci.ClientsImages.Validator()
}

func (oci OnlineClientsImages) GetNewOrDefaultImages() ([]byte, error) {
	log.Debug("fetching online client images")
	resp, err := http.Get(OnlineImagesSourceFile)
	if err != nil {
		log.Debugf("error fetching online images: %v", err)
		return []byte(clientImages), nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Debugf("error reading online images: %v", err)
		return []byte(clientImages), nil
	}

	return body, nil
}

func NewOnlineClientsImages() (ClientsImages, error) {
	oci := OnlineClientsImages{}

	rawImages, err := oci.GetNewOrDefaultImages()
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(rawImages, &oci.ClientsImages); err != nil {
		return nil, err
	}

	return &oci, nil
}
