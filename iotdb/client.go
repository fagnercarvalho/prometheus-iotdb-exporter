package iotdb

import (
	"github.com/apache/iotdb-client-go/client"
)

// Client represents a client for an IoTDB server instance.
type Client struct {
	Host        string
	Port        string
	Username    string
	Password    string
	TimeoutInMs int
}

// CountStorageGroups returns the storage groups count from a IoTDB instance.
func (c Client) CountStorageGroups() (int, error) {
	session, err := c.newSession()

	if err != nil {
		return -1, err
	}

	defer session.Close()

	resp, err := session.ExecuteQueryStatement("SHOW STORAGE GROUP")
	if err != nil {
		return -1, err
	}

	var count = 0
	for {
		hasNext, err := resp.Next()
		if err != nil {
			return -1, err
		}

		if !hasNext {
			return count, nil
		}

		count++
	}

	return -1, nil
}

// CountTimeSeries returns the time series count from a IoTDB instance.
func (c Client) CountTimeSeries() (int32, error) {
	session, err := c.newSession()

	if err != nil {
		return -1, err
	}

	defer session.Close()

	resp, err := session.ExecuteQueryStatement("COUNT TIMESERIES")
	if err != nil {
		return -1, err
	}

	hasNext, err := resp.Next()
	if err != nil {
		return -1, err
	}

	var count int32
	if !hasNext {
		return count, nil
	}

	if err := resp.Scan(&count); err != nil {
		return -1, err
	}

	return count, nil
}

// GetWriteAheadLogFileSize returns the write ahead log file size (in bytes) from a IoTDB instance.
func (c Client) GetWriteAheadLogFileSize() (int64, error) {
	session, err := c.newSession()

	if err != nil {
		return -1, err
	}

	defer session.Close()

	resp, err := session.ExecuteQueryStatement("SELECT LAST_VALUE(WAL) FROM root.stats.file_size")
	if err != nil {
		return -1, err
	}

	hasNext, err := resp.Next()
	if err != nil {
		return -1, err
	}

	var count int64
	if !hasNext {
		return count, nil
	}

	if err := resp.Scan(&count); err != nil {
		return -1, err
	}

	return count, nil
}

// GetSystemFileSize returns the system file size (in bytes) from a IoTDB instance.
func (c Client) GetSystemFileSize() (int64, error) {
	session, err := c.newSession()

	if err != nil {
		return -1, err
	}

	defer session.Close()

	resp, err := session.ExecuteQueryStatement("SELECT LAST_VALUE(SYS) FROM root.stats.file_size")
	if err != nil {
		return -1, err
	}

	hasNext, err := resp.Next()
	if err != nil {
		return -1, err
	}

	var count int64
	if !hasNext {
		return count, nil
	}

	if err := resp.Scan(&count); err != nil {
		return -1, err
	}

	return count, nil
}

// CountUsers returns the users count from a IoTDB instance.
func (c Client) CountUsers() (int, error) {
	session, err := c.newSession()

	if err != nil {
		return -1, err
	}

	defer session.Close()

	resp, err := session.ExecuteQueryStatement("LIST USER")
	if err != nil {
		return -1, err
	}

	var count = 0
	for {
		hasNext, err := resp.Next()
		if err != nil {
			return -1, err
		}

		if !hasNext {
			return count, nil
		}

		count++
	}

	return -1, nil
}

// PingServer pings a IoTDB instance to know if the server is running or not.
func (c Client) PingServer() error {
	session, err := c.newSession()

	if err != nil {
		return err
	}

	defer session.Close()

	return nil
}

func (c Client) newSession() (*client.Session, error) {
	session := client.NewSession(&client.Config{
		Host:     c.Host,
		Port:     c.Port,
		UserName: c.Username,
		Password: c.Password,
	})

	err := session.Open(false, c.TimeoutInMs)
	if err != nil {
		return nil, err
	}

	return session, nil
}

// NewClient instantiates a new client for an IoTDB server instance.
func NewClient(host string, port string, username string, password string, timeoutInMs int) Client {
	return Client{
		Host:        host,
		Port:        port,
		Username:    username,
		Password:    password,
		TimeoutInMs: timeoutInMs,
	}
}
