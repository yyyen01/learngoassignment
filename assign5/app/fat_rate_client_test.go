package app

import (
	"assignment/learngoassignment/assign5/pkg/apis"
	"testing"
	"time"
)

func TestWithBubleSort(t *testing.T) {
	client := FatRateClient{
		outputFilePath:   "../output/registrants.json",
		timeout:          2 * time.Second,
		totalRegistrants: 1000,
		sorter:           apis.BubbleSort{},
		writer:           &apis.JsonWriter{},
	}
	client.execute()
}

func TestWithQuickSort(t *testing.T) {
	client := FatRateClient{
		outputFilePath:   "../output/registrants.json",
		timeout:          2 * time.Second,
		totalRegistrants: 1000,
		sorter:           apis.QuickSort{},
		writer:           &apis.JsonWriter{},
	}
	client.execute()
}
