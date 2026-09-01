package metrics

import (
	"testing"

	dto "github.com/prometheus/client_model/go"
	"github.com/qdm12/gluetun/internal/configuration/settings"
	"github.com/qdm12/log"
	"github.com/stretchr/testify/assert"
)

type stubGatherer struct{}

func (stubGatherer) Gather() ([]*dto.MetricFamily, error) {
	return nil, nil
}

func Test_New(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		settings settings.Metrics
		expected string
	}{
		"noop type": {
			settings: settings.Metrics{
				Type: "noop",
			},
			expected: "noop metrics service",
		},
		"prometheus type": {
			settings: settings.Metrics{
				Type: "prometheus",
				Prometheus: settings.Prometheus{
					ListeningAddress: "127.0.0.1:0",
				},
			},
			expected: "prometheus http server",
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			service, err := New(testCase.settings, log.New(), stubGatherer{})
			assert.NoError(t, err)
			assert.Equal(t, testCase.expected, service.String())
		})
	}
}

func Test_New_UnknownTypePanics(t *testing.T) {
	t.Parallel()

	assert.PanicsWithValue(t, "unknown metrics type: unknown", func() {
		_, _ = New(settings.Metrics{Type: "unknown"}, log.New(), stubGatherer{})
	})
}
