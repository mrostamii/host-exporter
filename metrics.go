package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	OpsHostTcpAvailable = promauto.NewCounter(prometheus.CounterOpts{
		Name: "host_tcp_available_total",
		Help: "The total number of sucessful request",
	})
)

var (
	OpsHostTcpUnavailable = promauto.NewCounter(prometheus.CounterOpts{
		Name: "host_tcp_unavailable_total",
		Help: "The total number of sucessful request",
	})
)
