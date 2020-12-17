package exporter

import (
	"github.com/fagnercarvalho/prometheus-iotdb-exporter/iotdb"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	timeSeriesCount = prometheus.NewDesc(prometheus.BuildFQName("iotdb", "", "timeSeriesCount"), "Time Series Count", nil, nil)
)

type timeSeries struct {
}

func (s timeSeries) Name() string {
	return "storageGroup"
}

func (s timeSeries) Scrape(client iotdb.Client, ch chan<- prometheus.Metric) error {
	count, err := client.CountTimeSeries()
	if err != nil {
		return err
	}

	ch <- prometheus.MustNewConstMetric(timeSeriesCount, prometheus.GaugeValue, float64(count))

	return nil
}
