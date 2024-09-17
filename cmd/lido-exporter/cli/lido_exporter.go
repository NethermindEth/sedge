/*
Copyright 2022 Nethermind

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cli

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"os/signal"
	"slices"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/NethermindEth/sedge/cmd/lido-exporter/metrics"
	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/lido/contracts"
	"github.com/NethermindEth/sedge/internal/lido/contracts/csmodule"
	"github.com/NethermindEth/sedge/internal/utils"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	nested "github.com/antonfisher/nested-logrus-formatter"
	log "github.com/sirupsen/logrus"
)

var logLevel string

func RootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lido-exporter",
		Short: "Lido Exporter exports Lido CSM metrics to Prometheus",
		Long:  `Lido Exporter exports Lido CSM metrics to Prometheus`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			logLevel = viper.GetString("log-level")
			level, err := log.ParseLevel(strings.ToLower(logLevel))
			if err != nil {
				log.WithField(configs.Component, "Logger Init").Error(err)
				return
			}
			log.SetLevel(level)
		},
		Run: run,
	}

	// Disable completion default cmd
	cmd.CompletionOptions.DisableDefaultCmd = true

	// Persistent flags
	cmd.PersistentFlags().String("node-operator-id", "", "Node Operator ID")
	cmd.PersistentFlags().String("reward-address", "", "Reward address of Node Operator. It is used to calculate Node Operator ID if not set")
	cmd.PersistentFlags().String("network", "holesky", "Network name")
	cmd.PersistentFlags().StringSlice("rpc-endpoints", nil, "List of Ethereum HTTP RPC endpoints")
	cmd.PersistentFlags().StringSlice("ws-endpoints", nil, "List of Ethereum WebSocket RPC endpoints")
	cmd.PersistentFlags().String("port", "8080", "Port where the metrics will be exported.")
	cmd.PersistentFlags().Duration("scrape-time", 10*time.Second, "Time interval for scraping metrics. Values should be in the format of 10s, 1m, 1h, etc.")
	cmd.PersistentFlags().StringVar(&logLevel, "log-level", "info", "Set Log Level, e.g panic, fatal, error, warn, warning, info, debug, trace")

	viper.BindPFlag("node-operator-id", cmd.PersistentFlags().Lookup("node-operator-id"))
	viper.BindPFlag("reward-address", cmd.PersistentFlags().Lookup("reward-address"))
	viper.BindPFlag("network", cmd.PersistentFlags().Lookup("network"))
	viper.BindPFlag("rpc-endpoints", cmd.PersistentFlags().Lookup("rpc-endpoints"))
	viper.BindPFlag("ws-endpoints", cmd.PersistentFlags().Lookup("ws-endpoints"))
	viper.BindPFlag("port", cmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("scrape-time", cmd.PersistentFlags().Lookup("scrape-time"))
	viper.BindPFlag("log-level", cmd.PersistentFlags().Lookup("log-level"))
	viper.SetEnvPrefix("LIDO_EXPORTER")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	return cmd
}

func run(cmd *cobra.Command, args []string) {
	nodeOperatorID := viper.GetString("node-operator-id")
	rewardAddress := viper.GetString("reward-address")
	if nodeOperatorID == "" && rewardAddress == "" {
		log.Fatal("Node Operator ID or Reward Address is required")
	}

	// Validate port
	port := viper.GetString("port")
	_, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalf("Invalid port: %s", port)
	}

	network := viper.GetString("network")
	if !slices.Contains(configs.NetworkSupported(), network) {
		log.Fatalf("Invalid network: %s", network)
	}

	var nodeOperatorIDBigInt *big.Int
	if nodeOperatorID != "" {
		var ok bool
		nodeOperatorIDBigInt, ok = new(big.Int).SetString(nodeOperatorID, 10)
		if !ok {
			log.Fatalf("Failed to convert Node Operator ID to big.Int: %s", nodeOperatorID)
		}
		if nodeOperatorIDBigInt.Sign() < 0 { // Check if the value is negative
			log.Fatalf("Node Operator ID cannot be negative: %s", nodeOperatorID)
		}
	} else {
		if !utils.IsAddress(rewardAddress) {
			log.Fatalf("Invalid reward address: %s", rewardAddress)
		}

		var err error
		nodeOperatorIDBigInt, err = csmodule.NodeID(network, rewardAddress)
		if err != nil {
			log.Fatalf("Failed to get Node Operator ID: %v", err)
		}
	}

	rpcEndpoints := viper.GetStringSlice("rpc-endpoints")
	wsEndpoints := viper.GetStringSlice("ws-endpoints")

	client, err := contracts.ConnectClient(network, false, rpcEndpoints...)
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum RPC: %v", err)
	}
	wsClient, err := contracts.ConnectClient(network, true, wsEndpoints...)
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum WebSocket: %v", err)
	}

	// Initialize metrics
	metrics.InitMetrics(nodeOperatorID, network)

	// Start the metrics server
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", viper.GetString("port")), nil))
	}()

	// Start collecting metrics
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go metrics.CollectMetrics(ctx, client, wsClient, nodeOperatorIDBigInt, network, viper.GetDuration("scrape-time"))

	// Wait for interrupt signal to gracefully shutdown the exporter
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
	log.Info("Shutting down Lido Exporter...")
}

func init() {
	log.SetFormatter(&nested.Formatter{
		HideKeys:        true,
		FieldsOrder:     []string{configs.Component},
		TimestampFormat: "2006-01-02 15:04:05 --",
	})

	level, err := log.ParseLevel(strings.ToLower("error"))
	if err != nil {
		log.WithField(configs.Component, "Logger Init").Error(err)
		return
	}
	log.SetLevel(level)
	log.WithField(configs.Component, "Logger Init").Infof("Log level: %+v", logLevel)
}
