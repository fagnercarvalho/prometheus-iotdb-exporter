package exporter

import (
	"github.com/fagnercarvalho/prometheus-iotdb-exporter/iotdb"
	"github.com/prometheus/client_golang/prometheus"
)

type Scraper interface {
	Name() string

	Scrape(client iotdb.Client, ch chan<- prometheus.Metric) error
}
