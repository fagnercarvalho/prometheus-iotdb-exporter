package iotdb

import (
	"fmt"
	"testing"
)

func TestCountStorageGroups(t *testing.T) {
	client := NewClient("127.0.0.1", "6667", "root", "root")

	count, err := client.CountStorageGroups()
	if err != nil {
		t.Fatal(fmt.Sprintf("error when getting storage group count: %v", err))
	}

	fmt.Println(count)
}

func TestCountTimeSeries(t *testing.T) {
	client := NewClient("127.0.0.1", "6667", "root", "root")

	count, err := client.CountTimeSeries()
	if err != nil {
		t.Fatal(fmt.Sprintf("error when getting timeseries count: %v", err))
	}

	fmt.Println(count)
}

func TestGetWriteAheadLogFileSize(t *testing.T) {
	client := NewClient("127.0.0.1", "6667", "root", "root")

	size, err := client.GetWriteAheadLogFileSize()
	if err != nil {
		t.Fatal(fmt.Sprintf("error when getting write ahead log file size: %v", err))
	}

	fmt.Println(size)
}

func TestGetSystemFileSize(t *testing.T) {
	client := NewClient("127.0.0.1", "6667", "root", "root")

	size, err := client.GetSystemFileSize()
	if err != nil {
		t.Fatal(fmt.Sprintf("error when getting system file size: %v", err))
	}

	fmt.Println(size)
}

func TestCountUsers(t *testing.T) {
	client := NewClient("127.0.0.1", "6667", "root", "root")

	size, err := client.CountUsers()
	if err != nil {
		t.Fatal(fmt.Sprintf("error when getting user count: %v", err))
	}

	fmt.Println(size)
}
