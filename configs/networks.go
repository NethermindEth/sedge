package configs

import "time"

const (
	EpochTimeETH = 7 * time.Minute
	EpochTimeGNO = 2 * time.Minute
)

func SupportsMEVBoost(network string) bool {
	out, ok := networkConfigs[network]
	return ok && out.SupportsMEVBoost
}
