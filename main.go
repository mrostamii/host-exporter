package main

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var conf = LoadConfig()

func main() {
	recordMetrics()

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(conf.metricPort, nil))
}

func recordMetrics() {
	go func() {
		for {
			Tcp(conf.host, conf.port)
			time.Sleep(time.Duration(conf.checkInterval) * time.Second)
		}
	}()
}

func Tcp(host, port string) {
	timeout := time.Duration(1 * time.Second)
	_, err := net.DialTimeout("udp", host+":"+port, timeout)
	if err != nil {
		// fmt.Printf("%s %s %s\n", host, "not responding", err.Error()) // uncomment if you need log
		OpsHostTcpUnavailable.Inc()
	} else {
		// fmt.Printf("%s %s %s\n", host, "responding on port:", port) // uncomment if you need log
		OpsHostTcpAvailable.Inc()
	}
}
