package exporter

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
	"sync"
)
import "github.com/fagnercarvalho/prometheus-iotdb-exporter/iotdb"

type Config struct {
	Host        string
	Port        string
	Username    string
	Password    string
	TimeoutInMs int
}

type exporter struct {
	scrapers     []Scraper
	config       Config
}

func New(config Config) *exporter {
	exp := &exporter{config: config}

	exp.scrapers = append(exp.scrapers, storageGroup{})
	exp.scrapers = append(exp.scrapers, timeSeries{})
	exp.scrapers = append(exp.scrapers, user{})
	exp.scrapers = append(exp.scrapers, fileSize{})

	return exp
}

func (e *exporter) Describe(ch chan<- *prometheus.Desc) {
	//
}

func (e *exporter) Collect(ch chan<- prometheus.Metric) {
	e.scrape(ch)
}

func (e *exporter) scrape(ch chan<- prometheus.Metric) {
	client := iotdb.NewClient(e.config.Host, e.config.Port, e.config.Username, e.config.Password, e.config.TimeoutInMs)
	if err := client.PingServer(); err != nil {
		log.WithError(err).Error("Error while trying to connect to IoTDB")
		ch <- prometheus.NewInvalidMetric(
			prometheus.NewDesc("iotdb_connection_error",
				"Error while trying to connect to IoTDB", nil, nil),
			err)
		return
	}

	var wg sync.WaitGroup
	defer wg.Wait()
	for _, scraper := range e.scrapers {
		wg.Add(1)
		go func(s Scraper, client iotdb.Client) {
			defer wg.Done()
			err := s.Scrape(client, ch)
			if err != nil {
				log.WithError(err).WithFields(log.Fields{
					"scraper": s.Name(),
				}).Error("Error when scraping from IoTDB")

				ch <- prometheus.NewInvalidMetric(
					prometheus.NewDesc("iotdb_scrape_error",
						fmt.Sprintf("Error when scraping %v from IoTDB", s.Name()), nil, nil),
					err)
			}
		}(scraper, client)
	}
}
