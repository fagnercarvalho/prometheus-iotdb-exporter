## IoTDB Server Exporter [![Github Actions](https://img.shields.io/github/workflow/status/fagnercarvalho/prometheus-iotdb-exporter/Go)](https://github.com/fagnercarvalho/prometheus-iotdb-exporter/actions?query=workflow%3AGoo)

[![Go Report Card](https://goreportcard.com/badge/github.com/fagnercarvalho/prometheus-iotdb-exporter)](https://goreportcard.com/report/github.com/fagnercarvalho/prometheus-iotdb-exporter) [![Go Reference](https://pkg.go.dev/badge/github.com/fagnercarvalho/prometheus-iotdb-exporter.svg)](https://pkg.go.dev/github.com/fagnercarvalho/prometheus-iotdb-exporter) [![Docker Pulls](https://img.shields.io/docker/pulls/fagner/prometheus-iotdb-exporter.svg?maxAge=604800)](https://hub.docker.com/r/fagner/prometheus-iotdb-exporter) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/fagnercarvalho/prometheus-iotdb-exporter) [![Codecov](https://img.shields.io/codecov/c/gh/fagnercarvalho/prometheus-iotdb-exporter)](https://codecov.io/gh/fagnercarvalho/prometheus-iotdb-exporter)

Prometheus exporter for IoTDB server metrics.

Supported version: 0.11.0.

##### Running (Linux)

```
export IOTDB_PASSWORD=root
./prometheus-iotdb-exporter <flags>
```

##### Running (Windows)

```
set IOTDB_PASSWORD=root
prometheus-iotdb-exporter.exe <flags>
```

##### Flags

| Name          | Description             | Default   |
| ------------- | ----------------------- | --------- |
| listenPort    | exporter listening port | 8092      |
| iotDBHost     | IoTDB server host       | 127.0.0.1 |
| iotDBPort     | IoTDB server port       | 6667      |
| iotDBUsername | IoTDB username          | root      |

For security reasons the server password needs to be set by using `IOTDB_PASSWORD` environment variable.

##### Metrics

| Name                  | Metric Name                   | Description                                                  |
| --------------------- | ----------------------------- | ------------------------------------------------------------ |
| Write Ahead File Size | iotdb_write_ahead_file_size_bytes | Write Ahead File Size (extracted from the root.stats.file_size.WAL time series) in bytes. For this metric to be collected the enable_stat_monitor option must be enabled in the /iotdb/conf/iotdb-engine.properties config file. To update the config file check the [server_example](/server_example) folder. |
| System File Size      | iotdb_system_file_size_bytes       | System File Size (extracted from the root.stats.file_size.SYS time series) in bytes. For this metric to be collected the enable_stat_monitor option must be enabled in the /iotdb/conf/iotdb-engine.properties config file. To update the config file check the [server_example](/server_example) folder. |
| Storage Group Count   | iotdb_storage_groups         | Storage group count                                          |
| Timeseries Count      | iotdb_time_series           | Timeseries (across all storage groups) count                 |
| Users Count           | iotdb_users                  | Database users count                                         |

##### Docker

You can deploy this exporter by using the [fagner/prometheus-iotdb-exporter](https://hub.docker.com/r/fagner/prometheus-iotdb-exporter/) Docker image.

```
docker pull fagner/prometheus-iotdb-exporter

docker run -d -p "2000:8092" --name iotdb-exporter fagner/prometheus-iotdb-exporter
```

Or clone the repo and run the following commands.

```
docker build -t iotdb-exporter .
docker run -d -p "2000:8092" --name iotdb-exporter iotdb-exporter
```
