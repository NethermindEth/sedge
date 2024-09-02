package utils

import (
	"github.com/NethermindEth/sedge/internal/monitoring/common"
)

func init() {
	err := common.SetMockAVSs()
	if err != nil {
		panic(err)
	}
}
