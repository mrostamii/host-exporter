# **host-exporter**
Prometheus host exporter working as an agent.  
Feel free to add machines with this exporter. You will ensure that a host is available from multiple nodes and would not accessible just from a single machine.

## config.yml
```
# the endpoint you want to check
host: example.com 

# the port of the endpoint to check
port: 80

# how long in second it try checking (if using more, high load may happen to endpoint)
check-interval: 15

# the /metrics default port
metric-port: 8080
```

# Metrics
```
# HELP host_tcp_available_total The total number of sucessful request
# TYPE host_tcp_available_total counter
host_tcp_available_total 0
# HELP host_tcp_unavailable_total The total number of sucessful request
# TYPE host_tcp_unavailable_total counter
host_tcp_unavailable_total 0

```

# To check metrics
```
curl -sS localhost:8080/metrics | grep host_
```

# TODO

- [x] Prometheus metric
- [x] Load confing from yaml
- [ ] Dockerize + compose
- [ ] Prometheus service
- [ ] Multiple host
- [ ] Grafana service