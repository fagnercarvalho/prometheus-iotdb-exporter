package exporter

import (
	"github.com/fagnercarvalho/prometheus-iotdb-exporter/iotdb"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	writeAheadFileSize = prometheus.NewDesc(prometheus.BuildFQName("iotdb", "fileSize", "writeAheadSize"), "Write Ahead File Size (extracted from the root.stats.file_size.WAL time series) in bytes. For this metric to be collected the enable_stat_monitor option must be enabled in the /iotdb/conf/iotdb-engine.properties conf file.", nil, nil)
	systemFileSize     = prometheus.NewDesc(prometheus.BuildFQName("iotdb", "fileSize", "systemSize"), "System File Size (extracted from the root.stats.file_size.SYS time series) in bytes. For this metric to be collected the enable_stat_monitor option must be enabled in the /iotdb/conf/iotdb-engine.properties conf file.", nil, nil)
)

type fileSize struct {
}

func (f fileSize) Name() string {
	return "fileSize"
}

func (f fileSize) Scrape(client iotdb.Client, ch chan<- prometheus.Metric) error {
	errors := make(chan error, 0)

	go func() {
		size, err := client.GetWriteAheadLogFileSize()
		if err != nil {
			errors <- err
			return
		}

		ch <- prometheus.MustNewConstMetric(writeAheadFileSize, prometheus.GaugeValue, float64(size))
		errors <- nil
	}()

	go func() {
		size, err := client.GetSystemFileSize()
		if err != nil {
			errors <- err
			return
		}

		ch <- prometheus.MustNewConstMetric(systemFileSize, prometheus.GaugeValue, float64(size))
		errors <- nil
	}()

	for i := 0; i < 2; i++ {
		select {
		case err := <-errors:
			if err != nil {
				return err
			}
		}
	}

	return nil
}
