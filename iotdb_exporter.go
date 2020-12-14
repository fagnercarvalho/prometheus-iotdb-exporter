package main

import (
	"flag"
	"fmt"
	"github.com/fagnercarvalho/prometheus-iotdb-exporter/exporter"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/version"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

var (
	listenPort = flag.String("listenPort", "8092", "exporter listening port")
	iotDBHost  = flag.String("iotDBHost", "127.0.0.1", "IoTDB server host")
	iotDBPort  = flag.String("iotDBPort", "6667", "IoTDB server port")
	iotDBUsername = flag.String("iotDBUsername", "root", "IoTDB username")
	iotDBPassword = flag.String("iotDBPassword", "root", "IoTDB password")
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.TextFormatter{ForceColors: true})
}

func main() {
	flag.Parse()

	config := exporter.Config{
		Host: *iotDBHost,
		Port: *iotDBPort,
		Username: *iotDBUsername,
		Password: *iotDBPassword,
	}
	exporter := exporter.New(config)

	registry := prometheus.NewRegistry()
	registry.MustRegister(exporter)
	registry.MustRegister(version.NewCollector("iotdb"))

	h := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})

	log.WithFields(log.Fields{
		"port": *listenPort,
	}).Info("Serving IoTDB exporter metrics page")

	metricsPath := "/metrics"

	homePage := []byte(`<html>
<head><title>IoTDB Exporter</title></head>
<body>
<h2>IoTDB Exporter</h2>
<p><a href='`+ metricsPath +`'>Metrics</a></p>
</body></html>
`)

	http.Handle(metricsPath, h)
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Write(homePage)
	})
	err := http.ListenAndServe(fmt.Sprintf(":%v", *listenPort), nil)
	if err != nil {
		log.WithError(err).Error("Error when serving IoTDB exporter")
	}
}
