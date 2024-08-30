package data

import (
	"github.com/NethermindEth/eigenlayer/internal/common"
)

func init() {
	err := common.SetMockAVSs()
	if err != nil {
		panic(err)
	}
}
