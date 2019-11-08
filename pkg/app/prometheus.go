package app

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// QPS
	QPS = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "douqu",
		Subsystem: "media_management",
		Name:      "request_count",
		Help:      "http request count",
	}, []string{"path", "country"})

	// 400
	Req400 = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "douqu",
		Subsystem: "media_management",
		Name:      "400_request_count",
		Help:      "400 http request count",
	}, []string{"path", "country"})
)
