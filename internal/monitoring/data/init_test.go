package data

import (
	"github.com/NethermindEth/sedge/internal/common"
)

func init() {
	err := common.SetMockAVSs()
	if err != nil {
		panic(err)
	}
}
