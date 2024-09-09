package cli

import (
	"io"
	"log"
	"net/http"

	"github.com/NethermindEth/sedge/cli/actions"
	lidoExp "github.com/NethermindEth/sedge/internal/app/lido_exporter"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	cobra "github.com/spf13/cobra"
)

const (
	rpcAddress      = "wss://ethereum-rpc.publicnode.com"          //rpc address (mainnet)
	stringAddress   = "0xdAC17F958D2ee523a2206206994597C13D831ec7" // address of smart contract as string (tether)
	// rpcAddress = "wss://ethereum-holesky-rpc.publicnode.com" //rpc address (holesky)
	// stringAddress       = "0xffddf7025410412deaa05e3e1ce68fe53208afcb" // address of smart contract as string (ExitRequests Lido in Holesky)
	// stringAddress  = "0x4562c3e63c2e586cD1651B958C22F88135aCAd4f" // address of smart contract as string (ELRewardsStealingPenaltyReported  Lido in Holesky)
	// apiKey              = "" // api key to etherscan
)

func LExporterCmd(sedgeAction actions.SedgeActions) *cobra.Command {
	var flags GenCmdFlags

	cmd := &cobra.Command{
		Use:     "LExporter",
		Short:   "Start fetching and exposing metrics",
		Long:    `same`,
		Args:    cobra.NoArgs,
		PreRunE: nil,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runLExporterCmd(cmd.OutOrStdout(), &flags, sedgeAction, nil)
		},
	}
	return cmd
}

func runLExporterCmd(out io.Writer, flags *GenCmdFlags, sedgeAction actions.SedgeActions, services []string) error {
	client, err := ethclient.Dial(rpcAddress)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	go lidoExp.MonitorEvents(client, stringAddress)

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":2112", nil))

	return nil
}
