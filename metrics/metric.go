package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	appMetadata = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "app_metadata",
		Help: "Version information about this binary",
		ConstLabels: map[string]string{
			"app_name": "library-service",
			"version":  "v0.1.0",
			"language": "go",
		},
	})
	registry = prometheus.NewRegistry()

	HttpRequestsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Count of all HTTP requests",
	}, []string{"api", "method", "status"})
	MetricRegistry = registry
)

func InitMetrics() {
	registry.MustRegister(appMetadata)
	registry.MustRegister(HttpRequestsTotal)
}
