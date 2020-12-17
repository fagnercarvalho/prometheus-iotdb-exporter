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
	error        prometheus.Gauge
	iotDBUp      prometheus.Gauge
	scrapeErrors *prometheus.CounterVec
	config       Config
}

func New(config Config) *exporter {
	exp := &exporter{config: config}

	exp.scrapers = append(exp.scrapers, storageGroup{})
	exp.scrapers = append(exp.scrapers, timeSeries{})
	exp.scrapers = append(exp.scrapers, user{})
	exp.scrapers = append(exp.scrapers, fileSize{})

	exp.error = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "iotdb",
		Subsystem: "exporter",
		Name:      "last_scrape_error",
		Help:      "Whether the last scrape of metrics from iotDB resulted in an error (1 for error, 0 for success).",
	})

	exp.iotDBUp = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "iotdb",
		Subsystem: "exporter",
		Name:      "up",
		Help:      "Whether the iotDb server is up.",
	})

	exp.scrapeErrors = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "iotdb",
		Subsystem: "exporter",
		Name:      "scrape_errors_total",
		Help:      "Total number of times an error occurred scraping a iotDB.",
	}, []string{"collector"})

	return exp
}

func (e *exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- e.error.Desc()
	ch <- e.iotDBUp.Desc()
	e.scrapeErrors.Describe(ch)
	ch <- storageGroupCount
}

func (e *exporter) Collect(ch chan<- prometheus.Metric) {
	e.scrape(ch)
	ch <- e.error
	ch <- e.iotDBUp
	e.scrapeErrors.Collect(ch)
}

func (e *exporter) scrape(ch chan<- prometheus.Metric) {
	client := iotdb.NewClient(e.config.Host, e.config.Port, e.config.Username, e.config.Username, e.config.TimeoutInMs)
	if err := client.PingServer(); err != nil {
		log.WithError(err).Error("Error when connecting to IoTDB")
		e.error.Set(1)
		e.iotDBUp.Set(0)
		return
	}

	e.error.Set(0)
	e.iotDBUp.Set(1)

	var wg sync.WaitGroup
	defer wg.Wait()
	for _, scraper := range e.scrapers {
		wg.Add(1)
		go func(s Scraper, client iotdb.Client) {
			label := fmt.Sprintf("collector.%v", s.Name())
			defer wg.Done()
			err := s.Scrape(client, ch)
			if err != nil {
				log.WithError(err).WithFields(log.Fields{
					"scraper": s.Name(),
				}).Error("Error when scraping from IoTDB")

				e.error.Set(1)
				e.scrapeErrors.WithLabelValues(label).Inc()
			}
		}(scraper, client)
	}
}
