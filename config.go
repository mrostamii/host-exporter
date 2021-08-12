package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	host, port, metricPort string
	checkInterval          int
}

func LoadConfig() (config Config) {
	c := viper.New()
	c.SetConfigFile("config.yml")
	err := c.ReadInConfig()
	if err != nil {
		fmt.Println("The 'config.yml' not found!")
		OpsHostTcpUnavailable.Inc()
		return
	}

	conf := Config{
		host:          c.GetString("host"),
		port:          c.GetString("port"),
		checkInterval: c.GetInt("check-interval"),
		metricPort:    ":" + c.GetString("metric-port")}

	return conf
}
