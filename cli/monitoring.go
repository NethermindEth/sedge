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
	"errors"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/NethermindEth/sedge/internal/common"
	"github.com/NethermindEth/sedge/internal/monitoring"
)

func MonitoringCmd(mgr MonitoringManager) *cobra.Command {
	cmd := cobra.Command{
		Use:   "monitoring [init|clean]",
		Short: "Manage the monitoring stack",
		Long:  "Manage the monitoring stack. Use 'init' to install and run, or 'clean' to stop and uninstall.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			switch args[0] {
			case "init":
				return InitMonitoring(true, true, mgr)
			case "clean":
				return CleanMonitoring(mgr)
			default:
				return fmt.Errorf("invalid argument: %s. Use 'init' or 'clean'", args[0])
			}
		},
	}
	return &cmd
}

// Init initializes the Monitoring Stack. If install is true, it will install the Monitoring Stack if it is not installed.
// If run is true, it will run the Monitoring Stack if it is not running.
func InitMonitoring(install, run bool, monitoringMgr MonitoringManager) error {
	// Check if the monitoring stack is installed.
	installStatus, err := monitoringMgr.InstallationStatus()
	if err != nil {
		return err
	}
	log.Debugf("Monitoring stack installation status: %v", installStatus == common.Installed)
	// If the monitoring stack is not installed, install it.
	if installStatus == common.NotInstalled && install {
		err = monitoringMgr.InstallStack()
		if errors.Is(err, monitoring.ErrInstallingMonitoringMngr) {
			// If the monitoring stack installation fails, remove the monitoring stack directory.
			if cerr := monitoringMgr.Cleanup(); cerr != nil {
				return fmt.Errorf("install failed: %w. Failed to cleanup monitoring stack after installation failure: %w", err, cerr)
			}
			return err
		} else if err != nil {
			return err
		}
	}
	// Check if the monitoring stack is running.
	status, err := monitoringMgr.Status()
	if err != nil {
		log.Debugf("Monitoring stack status: unknown. Got error: %v", err)
	}
	// If the monitoring stack is not running, start it.
	if status != common.Running && status != common.Restarting && run {
		if err := monitoringMgr.Run(); err != nil {
			return err
		}
	} else if status != common.Running && status != common.Restarting && !run {
		// If the monitoring stack is not supposed to be running then exit.
		return nil
	}

	// Initialize monitoring stack if it is running.
	if err := monitoringMgr.Init(); err != nil {
		return err
	}

	return nil
}

// CleanMonitoring stops and uninstalls the Monitoring Stack
func CleanMonitoring(monitoringMgr MonitoringManager) error {
	// Check if the monitoring stack is installed.
	installStatus, err := monitoringMgr.InstallationStatus()
	if err != nil {
		return err
	}
	log.Debugf("Monitoring stack installation status: %v", installStatus == common.Installed)
	// If the monitoring stack is installed, uninstall it.
	if installStatus == common.Installed {
		if err := monitoringMgr.Cleanup(); err != nil {
			return err
		}
	}
	return nil
}
