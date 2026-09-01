// Package metrics sets up metrics and creates a metrics service.
package metrics

import (
	dto "github.com/prometheus/client_model/go"
	"github.com/qdm12/gluetun/internal/configuration/settings"
	"github.com/qdm12/gluetun/internal/metrics/noop"
	"github.com/qdm12/gluetun/internal/metrics/prometheus"
	"github.com/qdm12/goservices"
	"github.com/qdm12/log"
)

// ParentLogger is the interface to create a new logger
// for the metrics service.
type ParentLogger interface {
	New(options ...log.Option) *log.Logger
}

// PromGatherer is the interface for gathering Prometheus metrics.
type PromGatherer interface {
	Gather() ([]*dto.MetricFamily, error)
}

// New creates a new metrics service based on the
// metrics type in the settings. It panics if the
// type is unknown, which should not happen if the
// settings were validated.
func New(settings settings.Metrics, parentLogger ParentLogger, //nolint:ireturn
	promGatherer PromGatherer,
) (service goservices.Service, err error) {
	switch settings.Type {
	case "noop":
		return noop.New()
	case "prometheus":
		logger := parentLogger.New(log.SetComponent("prometheus server"))
		return prometheus.New(settings.Prometheus, promGatherer, logger)
	default:
		panic("unknown metrics type: " + settings.Type)
	}
}
