package exporter

import (
	"github.com/fagnercarvalho/prometheus-iotdb-exporter/iotdb"
	"github.com/prometheus/client_golang/prometheus"
)

// Scraper represents a scraper that gets a metric from a IoTDB instance.
type Scraper interface {
	Name() string

	Scrape(client iotdb.Client, ch chan<- prometheus.Metric) error
}
