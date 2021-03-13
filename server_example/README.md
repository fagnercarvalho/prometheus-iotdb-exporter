### IoTDB server example with config changes

This is an example of an IoTDB server instance running with a change in the `iotdb-engine.properties.conf` file.

To run this example, run the following commands:

```
docker build -t iotdb .
docker run -d -p 0.0.0.0:6667:6667 -p 0.0.0.0:31999:31999 -p 0.0.0.0:8181:8181 --name example-iotdb iotdb
```

Then access the CLI and confirm the `root.stats` storage group was created.

```
docker exec -it example-iotdb /bin/bash
start-cli.sh
SHOW TIMESERIES
```

