package iotdb

import (
	"github.com/manlge/go-iotdb/client"
)

type Client struct {
	Host        string
	Port        string
	Username    string
	Password    string
	TimeoutInMs int
}

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

func (c Client) PingServer() error {
	session, err := c.newSession()

	if err != nil {
		return err
	}

	defer session.Close()

	return nil
}

func (c Client) newSession() (client.Session, error) {
	session := client.NewSession(c.Host, c.Port)

	session.User = c.Username
	session.Passwd = c.Password

	err := session.Open(false, c.TimeoutInMs)
	if err != nil {
		return client.Session{}, err
	}

	return session, nil
}

func NewClient(host string, port string, username string, password string, timeoutInMs int) Client {
	return Client{host, port, username, password, timeoutInMs}
}
