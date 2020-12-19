package exporter

import (
	"github.com/fagnercarvalho/prometheus-iotdb-exporter/iotdb"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	userCount = prometheus.NewDesc(prometheus.BuildFQName("iotdb", "", "user_count"), "User Count", nil, nil)
)

type user struct {
}

func (u user) Name() string {
	return "user"
}

func (u user) Scrape(client iotdb.Client, ch chan<- prometheus.Metric) error {
	count, err := client.CountUsers()
	if err != nil {
		return err
	}

	ch <- prometheus.MustNewConstMetric(userCount, prometheus.GaugeValue, float64(count))

	return nil
}
