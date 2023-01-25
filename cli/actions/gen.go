package actions

//go:generate mockgen -package=sedge_mocks -destination=../../mocks/sedgeActions.go github.com/NethermindEth/sedge/cli/actions SedgeActions

//go:generate mockgen -package=sedge_mocks -destination=../../mocks/depsHandler.go github.com/NethermindEth/sedge/cli/actions DependenciesHandlers
