package metrics

import (
	"context"

	prometheus "github.com/prometheus/client_golang/prometheus"
	promauto "github.com/prometheus/client_golang/prometheus/promauto"

	consts "github.com/MGomed/auth/consts"
)

// Metrics is api metrics
type Metrics struct {
	requestCounter        prometheus.Counter
	responseCounter       *prometheus.CounterVec
	histogramResponseTime *prometheus.HistogramVec
}

var metrics *Metrics

// Init inits collected metrics
func Init(_ context.Context) error {
	metrics = &Metrics{
		requestCounter: promauto.NewCounter(
			prometheus.CounterOpts{
				Namespace: consts.MetricsNamespace,
				Subsystem: "grpc",
				Name:      consts.MetricsAppName + "_requests_total",
				Help:      "Server's request count",
			},
		),
		responseCounter: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: consts.MetricsNamespace,
				Subsystem: "grpc",
				Name:      consts.MetricsAppName + "_responses_total",
				Help:      "Server's response count",
			},
			[]string{consts.MetricStatusLabel, consts.MetricMethodLabel},
		),
		histogramResponseTime: promauto.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: consts.MetricsNamespace,
				Subsystem: "grpc",
				Name:      consts.MetricsAppName + "_histogram_response_time_seconds",
				Help:      "Server's response time",
				Buckets:   prometheus.ExponentialBuckets(0.0001, 2, 16),
			},
			[]string{consts.MetricStatusLabel},
		),
	}

	return nil
}

// IncRequestCounter is request count metric
func IncRequestCounter() {
	metrics.requestCounter.Inc()
}

// IncResponseCounter is response count metric
func IncResponseCounter(status string, method string) {
	metrics.responseCounter.WithLabelValues(status, method).Inc()
}

// HistogramResponseTimeObserve is response time metric
func HistogramResponseTimeObserve(status string, time float64) {
	metrics.histogramResponseTime.WithLabelValues(status).Observe(time)
}
