package storage

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"testing"
	"timekeeper/models"
)

func getTestFilePath(filename string) (string, error) {
	tmpDir := os.TempDir()
	return filepath.Join(tmpDir, "timekeeper_test", filename), nil
}

func TestSaveEntryToCSV(t *testing.T) {
	originalGetConfigPath := getConfigPath
	getConfigPath = getTestFilePath
	defer func() { getConfigPath = originalGetConfigPath }()

	entry := models.Entry{
		Customer: "TestCorp",
		Task:     "Write test",
		Minutes:  90,
		Date:     "2025-04-11",
	}

	err := SaveEntryToCSV(entry)
	if err != nil {
		t.Fatalf("SaveEntryToCSV failed: %v", err)
	}

	path, _ := getConfigPath("entries.csv")
	file, err := os.Open(path)
	if err != nil {
		t.Fatalf("Could not open file: %v", err)
	}
	defer file.Close()
	defer os.Remove(path)

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		t.Fatalf("Reading CSV failed: %v", err)
	}

	if len(records) == 0 || records[0][0] != "TestCorp" {
		t.Errorf("Expected to find 'TestCorp', got: %v", records)
	}
}

func TestSaveCustomerToCSV(t *testing.T) {
	originalGetConfigPath := getConfigPath
	getConfigPath = getTestFilePath
	defer func() { getConfigPath = originalGetConfigPath }()

	customer := models.Customer{Name: "TestClient"}

	err := SaveCustomerToCSV(customer)
	if err != nil {
		t.Fatalf("SaveCustomerToCSV failed: %v", err)
	}

	path, _ := getConfigPath("customers.csv")
	file, err := os.Open(path)
	if err != nil {
		t.Fatalf("Could not open file: %v", err)
	}
	defer file.Close()
	defer os.Remove(path)

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		t.Fatalf("Reading CSV failed: %v", err)
	}

	if len(records) == 0 || records[0][0] != "TestClient" {
		t.Errorf("Expected to find 'TestClient', got: %v", records)
	}
}
