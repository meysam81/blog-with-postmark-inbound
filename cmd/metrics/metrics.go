package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	HTTPRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "tarzan_http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path", "status_code"},
	)

	HTTPRequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "tarzan_http_request_duration_seconds",
			Help:    "HTTP request duration in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path"},
	)

	WebSocketConnections = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "tarzan_websocket_connections_active",
			Help: "Number of active WebSocket connections",
		},
	)

	WebSocketConnectionsTotal = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "tarzan_websocket_connections_total",
			Help: "Total number of WebSocket connections",
		},
	)

	PostsTotal = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "tarzan_posts_total",
			Help: "Total number of posts created",
		},
	)

	PostsStorageBytes = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "tarzan_posts_storage_bytes",
			Help: "Total storage used by posts in bytes",
		},
	)

	EmailsProcessedTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "tarzan_emails_processed_total",
			Help: "Total number of emails processed",
		},
		[]string{"status"},
	)

	EmailProcessingDuration = promauto.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "tarzan_email_processing_duration_seconds",
			Help:    "Email processing duration in seconds",
			Buckets: prometheus.DefBuckets,
		},
	)

	DatabaseConnectionsActive = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "tarzan_database_connections_active",
			Help: "Number of active database connections",
		},
	)

	DatabaseOperationsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "tarzan_database_operations_total",
			Help: "Total number of database operations",
		},
		[]string{"operation", "status"},
	)

	DatabaseOperationDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "tarzan_database_operation_duration_seconds",
			Help:    "Database operation duration in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"operation"},
	)
)

func RecordHTTPRequest(method, path, statusCode string) {
	HTTPRequestsTotal.WithLabelValues(method, path, statusCode).Inc()
}

func RecordHTTPDuration(method, path string, duration float64) {
	HTTPRequestDuration.WithLabelValues(method, path).Observe(duration)
}

func IncrementWebSocketConnections() {
	WebSocketConnections.Inc()
	WebSocketConnectionsTotal.Inc()
}

func DecrementWebSocketConnections() {
	WebSocketConnections.Dec()
}

func IncrementPostsTotal() {
	PostsTotal.Inc()
}

func UpdatePostsStorageBytes(bytes float64) {
	PostsStorageBytes.Set(bytes)
}

func RecordEmailProcessed(status string) {
	EmailsProcessedTotal.WithLabelValues(status).Inc()
}

func RecordEmailProcessingDuration(duration float64) {
	EmailProcessingDuration.Observe(duration)
}

func UpdateDatabaseConnections(count float64) {
	DatabaseConnectionsActive.Set(count)
}

func RecordDatabaseOperation(operation, status string) {
	DatabaseOperationsTotal.WithLabelValues(operation, status).Inc()
}

func RecordDatabaseOperationDuration(operation string, duration float64) {
	DatabaseOperationDuration.WithLabelValues(operation).Observe(duration)
}
