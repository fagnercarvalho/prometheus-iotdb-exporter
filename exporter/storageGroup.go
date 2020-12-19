package exporter

import (
	"github.com/fagnercarvalho/prometheus-iotdb-exporter/iotdb"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	storageGroupCount = prometheus.NewDesc(prometheus.BuildFQName("iotdb", "", "storage_groups"), "Storage Group Count", nil, nil)
)

type storageGroup struct {
}

func (s storageGroup) Name() string {
	return "storageGroup"
}

func (s storageGroup) Scrape(client iotdb.Client, ch chan<- prometheus.Metric) error {
	count, err := client.CountStorageGroups()
	if err != nil {
		return err
	}

	ch <- prometheus.MustNewConstMetric(storageGroupCount, prometheus.GaugeValue, float64(count))

	return nil
}
