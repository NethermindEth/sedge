package monitoring

import "errors"

var (
	ErrInitializingMonitoringMngr    = errors.New("error initializing monitoring manager")
	ErrCheckingMonitoringStack       = errors.New("error checking monitoring stack status")
	ErrRunningMonitoringStack        = errors.New("error running monitoring stack")
	ErrInstallingMonitoringMngr      = errors.New("error installing monitoring manager")
	ErrConfiguringMonitoringServices = errors.New("error configuring monitoring services")
	ErrNonexistingTarget             = errors.New("target to remove does not exist")
)
