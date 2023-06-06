package clientsimages

import (
	"io"
	"net/http"
	"os"
	"path"

	"github.com/NethermindEth/sedge/configs"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

const (
	OnlineImagesSourceFile = "https://raw.githubusercontent.com/NethermindEth/sedge/develop/configs/client_images.yaml"
	OnlineImagesCacheFile  = "sedge_client_images.yaml"
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

func (oci OnlineClientsImages) getCachedImages() ([]byte, error) {
	return os.ReadFile(oci.getCachedImagesFilePath())
}

func (oci OnlineClientsImages) setCachedImages() error {
	body, err := yaml.Marshal(oci.ClientsImages)
	if err != nil {
		return err
	}
	return os.WriteFile(oci.getCachedImagesFilePath(), body, 0666)
}

func (oci OnlineClientsImages) getCachedImagesFilePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		home = "."
	}

	return path.Join(home, configs.ConfigFolderName, OnlineImagesCacheFile)
}

func (oci OnlineClientsImages) GetNewOrDefaultImages() ([]byte, error) {
	var rawImages []byte
	log.Debug("fetching online client images")
	resp, err := http.Get(OnlineImagesSourceFile)
	if err != nil {
		log.Debugf("error fetching online images: %v", err)
		rawImages, err = oci.getCachedImages()
		if err != nil {
			log.Debugf("error getting cached images: %v", err)
			rawImages = []byte(clientImages)
		}
	} else {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Debugf("error reading online images: %v", err)
			rawImages, err = oci.getCachedImages()
			if err != nil {
				log.Debugf("error getting cached images: %v", err)
				rawImages = []byte(clientImages)
			}
		} else {
			rawImages = body
		}
	}

	return rawImages, nil
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

	if err := oci.setCachedImages(); err != nil {
		log.Debugf("error setting cached images: %v", err)
	}

	return &oci, nil
}
