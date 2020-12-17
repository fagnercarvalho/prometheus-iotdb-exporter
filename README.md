## IoTDB Server Exporter

Prometheus exporter for IoTDB server metrics.

Supported version: 0.11.0.

##### Running

```
./iotdb-exporter <flags>
```

##### Flags

| Name          | Description             | Default   |
| ------------- | ----------------------- | --------- |
| listenPort    | exporter listening port | 8092      |
| iotDBHost     | IoTDB server host       | 127.0.0.1 |
| iotDBPort     | IoTDB server port       | 6667      |
| iotDBUsername | IoTDB username          | root      |
| iotDPassword  | IoTDB password          | root      |

##### Metrics

| Name                  | Description                                                  |
| --------------------- | ------------------------------------------------------------ |
| Write Ahead File Size | Write Ahead File Size (extracted from the root.stats.file_size.WAL time series) in bytes. For this metric to be collected the enable_stat_monitor option must be enabled in the /iotdb/conf/iotdb-engine.properties config file. To update the config file check the [server_example](/server_example) folder. |
| System File Size      | System File Size (extracted from the root.stats.file_size.SYS time series) in bytes. For this metric to be collected the enable_stat_monitor option must be enabled in the /iotdb/conf/iotdb-engine.properties config file. To update the config file check the [server_example](/server_example) folder. |
| Storage Group Count   | Storage group count                                          |
| Timeseries Count      | Timeseries (across all storage groups) count                 |
| Users Count           | Database users count                                         |

##### Docker

Clone the repo and run the following commands.

```
docker build -t iotdb-exporter .
docker run -d -p "2000:8092" --name iotdb-exporter iotdb-exporter
```

