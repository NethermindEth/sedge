package monitoring

//go:generate mockgen -package=sedge_mocks -destination=../../mocks/monitoring_service.go -package=mocks github.com/NethermindEth/sedge/internal/pkg/monitoring ServiceAPI

//go:generate mockgen -package=sedge_mocks -destination=../../mocks/compose.go -package=mocks github.com/NethermindEth/sedge/internal/pkg/monitoring ComposeManager

//go:generate mockgen -package=sedge_mocks -destination=../../mocks/docker.go -package=mocks github.com/NethermindEth/sedge/internal/pkg/monitoring DockerServiManager
