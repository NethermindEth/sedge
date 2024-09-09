package lido_exporter

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	eventCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "blockchain_event_count",
			Help: "Count of blockchain events.",
		},
		[]string{"event_type"},
	)
)

func init() {
	prometheus.MustRegister(eventCounter)
}
