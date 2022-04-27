package configs_test

import (
	"fmt"
	"testing"

	"github.com/NethermindEth/1Click/configs"
)

func TestGetConfigClients(t *testing.T) {
	data := configs.GetConfigClients

	for d := range data["validator"] {
		fmt.Println(d)
	}
}
